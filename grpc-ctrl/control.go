/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: control.go
 * @time: 2019/8/12 15:07
 */
package grpc_ctrl

import (
	`context`
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	`sync`
	`time`

	`github.com/astaxie/beego/logs`
	gp `github.com/flyaways/pool`
	libs `github.com/generalzgd/comm-libs`
	`github.com/generalzgd/comm-libs/conf/ymlcfg`
	`github.com/generalzgd/deepcopy/dcopy`

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	`google.golang.org/grpc/metadata`
	`google.golang.org/grpc/peer`

	`github.com/generalzgd/grpc-svr-frame/common`
)

type GrpcController struct {
	connLock    sync.Mutex
	grpcConnMap map[string]*gp.GRPCPool
}

func MakeGrpcController() GrpcController {
	return GrpcController{
		grpcConnMap: map[string]*gp.GRPCPool{},
	}
}

func (p *GrpcController) GetPeer(ctx context.Context) (*peer.Peer, bool) {
	pe, ok := peer.FromContext(ctx)
	return pe, ok
}

// get client info from in coming context
func (p *GrpcController) GetClientInfo(ctx context.Context) (*common.ClientConnInfo, metadata.MD) {
	connInfo := &common.ClientConnInfo{}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		m := p.Md2MapInterface(md)
		dcopy.InstanceFromMap(&connInfo, m)
	}
	// logs.Debug("connInfo value: %v", connInfo)

	return connInfo, md
}

// make out going context by metadata
func (p *GrpcController) MakeOutgoingContext(ctx context.Context, md metadata.MD) context.Context {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

func (p *GrpcController) MakeIncomingContext(ctx context.Context, md metadata.MD) context.Context {
	ctx = metadata.NewIncomingContext(ctx, md)
	return ctx
}

// make out going context by client conn info
func (p *GrpcController) MakeOutgoingContextByClientInfo(ctx context.Context, info *common.ClientConnInfo) context.Context {
	md, _ := p.ClientInfoToMD(info)
	return p.MakeOutgoingContext(ctx, md)
}

func (p *GrpcController) MakeIncomingContextByClientInfo(ctx context.Context, info *common.ClientConnInfo) context.Context {
	md, _ := p.ClientInfoToMD(info)
	return p.MakeIncomingContext(ctx, md)
}

func (p *GrpcController) Md2MapInterface(md metadata.MD) map[string]interface{} {
	out := make(map[string]interface{}, len(md))
	for k, v := range md {
		if len(v) > 0 {
			out[k] = v[0]
		}
	}
	return out
}

func (p *GrpcController) ClientInfoToMD(info *common.ClientConnInfo) (metadata.MD, error) {
	m, err := dcopy.InstanceToMap(info)
	if err != nil {
		return nil, err
	}
	md := metadata.New(libs.MapInterface2String(m))
	return md, nil
}

// 获取tls配置信息
func (p *GrpcController) GetCreds(svrName string, cfg ymlcfg.CertFile) (grpc.DialOption, error) {
	certificate, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Priv)
	if err != nil {
		logs.Error("load cert file fail.", err)
		return nil, err
	}
	certPool := x509.NewCertPool()
	f, err := os.Open(cfg.RootCAFile)
	if err != nil {
		logs.Error("load ca file fail.", err)
		return nil, err
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		return nil, err
	}
	transportCreds := credentials.NewTLS(&tls.Config{
		ServerName:   svrName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})
	otp := grpc.WithTransportCredentials(transportCreds)
	return otp, nil
}

// 获取grpc Dial option
func (p *GrpcController) GetDialOption(cfg ymlcfg.EndpointConfig) []grpc.DialOption {
	opts := []grpc.DialOption{
		grpc.WithReadBufferSize(10 * 1024),
		grpc.WithWriteBufferSize(10 * 1024),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(32*1024), grpc.MaxCallSendMsgSize(32*1024)),
	}
	if cfg.Secure {
		opt, _ := p.GetCreds(cfg.Name, cfg.CertFiles[0])
		opts = append(opts, opt)
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	return opts
}

func (p *GrpcController) GetTlsConfig(certFiles ...ymlcfg.CertFile) (*tls.Config, error) {
	if len(certFiles) < 1 {
		return nil, errors.New("cert file empty")
	}
	tlsCfg := &tls.Config{}
	tlsCfg.MaxVersion = tls.VersionTLS10
	tlsCfg.Certificates = make([]tls.Certificate, len(certFiles))
	dir := filepath.Dir(os.Args[0])
	var err error
	for i, certFile := range certFiles {
		certFilePath := filepath.Join(dir, "certs", certFile.Cert)
		privFilePath := filepath.Join(dir, "certs", certFile.Priv)
		if tlsCfg.Certificates[i], err = tls.LoadX509KeyPair(certFilePath, privFilePath); err != nil {
			logs.Error("load cert file fail.", err)
			return nil, err
		}
	}
	return tlsCfg, nil
}

// 单纯获取客户端链接
func (p *GrpcController) GetGrpcConn(key, addr string, cfg ymlcfg.EndpointConfig) (*grpc.ClientConn, func(), error) {
	p.connLock.Lock()
	defer p.connLock.Unlock()

	// 服务名
	if pool, ok := p.grpcConnMap[key]; ok {
		conn, err := pool.Get()
		if err != nil {
			return nil, nil, err
		}
		return conn, func() {
			pool.Put(conn)
		}, nil
	}

	options := &gp.Options{
		InitTargets:  []string{addr},
		InitCap:      1,
		MaxCap:       10,
		DialTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	pool, err := gp.NewGRPCPool(options, p.GetDialOption(cfg)...) // , cfg.Pool.InitNum, cfg.Pool.CapNum, cfg.Pool.IdleTimeout, cfg.Pool.LifeDur)
	if err != nil {
		return nil, nil, err
	}
	p.grpcConnMap[key] = pool
	if conn, err := pool.Get(); err != nil {
		return nil, nil, err
	} else {
		return conn, func() {
			pool.Put(conn)
		}, nil
	}
}

// 获取rpc客户端链接，包含round-robin策略
func (p *GrpcController) GetGrpcConnWithLB(cfg ymlcfg.EndpointConfig) (*grpc.ClientConn, func(), error) {
	p.connLock.Lock()
	defer p.connLock.Unlock()

	// 服务名
	if pool, ok := p.grpcConnMap[cfg.Name]; ok {
		conn, err := pool.Get()
		if err != nil {
			return nil, nil, err
		}
		return conn, func() {
			pool.Put(conn)
		}, nil
	}

	options := &gp.Options{
		InitTargets:  []string{cfg.Address},
		InitCap:      1,
		MaxCap:       10,
		DialTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	pool, err := gp.NewGRPCPool(options, p.GetDialOption(cfg)...) // , cfg.Pool.InitNum, cfg.Pool.CapNum, cfg.Pool.IdleTimeout, cfg.Pool.LifeDur)
	if err != nil {
		return nil, nil, err
	}
	p.grpcConnMap[cfg.Name] = pool
	if conn, err := pool.Get();err != nil {
		return nil, nil, err
	} else {
		return conn, func() {
			pool.Put(conn)
		}, nil
	}
}

func (p *GrpcController) GetGrpcConnWithLBValues(name string, address string, port int, secure bool, cert string, priv string, rootCaFile string) (*grpc.ClientConn, func(), error) {
	cfg := ymlcfg.EndpointConfig{
		Name:    name,
		Address: address,
		Port:    port,
		Secure:  secure,
		CertFiles: []ymlcfg.CertFile{
			{
				Cert:       cert,
				Priv:       priv,
				RootCAFile: rootCaFile,
			},
		},
	}

	return p.GetGrpcConnWithLB(cfg)
}

// 销毁所有链接, 或指定销毁
func (p *GrpcController) DisposeGrpcConn(svrnameOrKey string) {
	p.connLock.Lock()
	defer p.connLock.Unlock()

	if svrnameOrKey != "" {
		if pool, ok := p.grpcConnMap[svrnameOrKey]; ok {
			pool.Close()
			delete(p.grpcConnMap, svrnameOrKey)
		}
		return
	}
	for _, pool := range p.grpcConnMap {
		pool.Close()
	}
	p.grpcConnMap = map[string]*gp.GRPCPool{}
}
