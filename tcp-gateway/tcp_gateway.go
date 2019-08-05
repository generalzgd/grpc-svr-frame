/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: tcp_gateway.go
 * @time: 2019-07-17 19:42
 */

package tcp_gateway

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/funny/link"
	"github.com/funny/slab"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"

	"github.com/generalzgd/grpc-svr-frame/common"
	"github.com/generalzgd/grpc-svr-frame/config"
	"github.com/generalzgd/grpc-svr-frame/grpclb"

	gate_msg "github.com/generalzgd/grpc-svr-frame/tcp-gateway/proto"
)

const (
	CONTENT_LENGTH_LIMIT         = 1024 * 32
	TCP_GATE_PACKAGE_HEAD_LENGTH = 8
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type GrpcSvrItem struct {
	grpclb.GrpcSvrInfo
	Conn *grpc.ClientConn
	Ok   bool // 服务健康状态
}

type TcpGate struct {
	GateProtocol
	config   config.GateConnConfig
	listener net.Listener
	//
	connSeed uint32
	server   *link.Server
	//
	closeOnce sync.Once
	workerNum uint16
	exit      chan struct{}
	closeFlag bool
}

func (p *TcpGate) Initialize(pool slab.Pool) {
	p.PackEncrypt.Init(TCP_GATE_PACKAGE_HEAD_LENGTH, CONTENT_LENGTH_LIMIT, pool)
	p.exit = make(chan struct{})
}

func (p *TcpGate) Destroy() {
	p.closeOnce.Do(func() {
		close(p.exit)
		p.closeFlag = true
	})
}

func (p *TcpGate) Serve(cfg config.GwItemConfig) error {
	switch cfg.Type {
	case common.GW_TYPE_HTTP:
	case common.GW_TYPE_TCP:
		return p.serveTcp(cfg.GwConf)
	case common.GW_TYPE_WS:
		return p.serveWs(cfg.GwConf)
	}
	return nil
}

func (p *TcpGate) serveTcp(cfg config.GateConnConfig) error {
	p.config = cfg
	addr := fmt.Sprintf(":%d", cfg.Port)

	if cfg.Insecure {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		p.server = link.NewServer(
			lis,
			link.ProtocolFunc(func(rw io.ReadWriter) (link.Codec, error) {
				return p.newTcpCodec(atomic.AddUint32(&p.connSeed, 1), rw.(net.Conn), cfg.BufferSize), nil
			}),
			cfg.SendChanSize,
			link.HandlerFunc(func(session *link.Session) {
				p.handleSession(session, cfg.MaxConn)
			}),
		)
	} else {
		if len(cfg.CertFiles) < 1 {
			return errors.New("cert files need")
		}
		tlsCfg := &tls.Config{}
		tlsCfg.MaxVersion = tls.VersionTLS11
		tlsCfg.Certificates = make([]tls.Certificate, len(cfg.CertFiles))

		dir := filepath.Dir(os.Args[0])
		var err error
		for i, certFile := range cfg.CertFiles {
			certFilePath := filepath.Join(dir, "certs", certFile.Cert)
			privFilePath := filepath.Join(dir, "certs", certFile.Priv)
			if tlsCfg.Certificates[i], err = tls.LoadX509KeyPair(certFilePath, privFilePath); err != nil {
				beego.Error("Load cert file failure.", err)
				return err
			}
		}

		ln, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		lis := tls.NewListener(ln, tlsCfg)

		p.server = link.NewServer(
			lis,
			link.ProtocolFunc(func(rw io.ReadWriter) (link.Codec, error) {
				return p.newTlsCodec(atomic.AddUint32(&p.connSeed, 1), rw.(net.Conn), cfg.BufferSize), nil
			}),
			cfg.SendChanSize,
			link.HandlerFunc(func(session *link.Session) {
				p.handleSession(session, cfg.MaxConn)
			}),
		)
	}

	go func() {
		if err := p.server.Serve(); err != nil {
			logs.Error("tcp gate way server error.", err)
		}
	}()

	logs.Info("start listen tcp with secure.", addr, cfg.Insecure)
	return nil
}

func (p *TcpGate) serveWs(cfg config.GateConnConfig) error {
	var (
		httpServeMux = http.NewServeMux()
		err          error
	)
	httpServeMux.HandleFunc("/", p.serveWebSocket)
	httpServeMux.HandleFunc("/health", p.serveWebHealth)
	addr := fmt.Sprintf(":%d", cfg.Port)
	tlsCfg := &tls.Config{}
	tlsCfg.Certificates = make([]tls.Certificate, len(cfg.CertFiles))
	bWss := !cfg.Insecure
	if bWss {
		dir := filepath.Dir(os.Args[0])
		for i, certFile := range cfg.CertFiles {
			certFilePath := filepath.Join(dir, "certs", certFile.Cert)
			privFilePath := filepath.Join(dir, "certs", certFile.Priv)
			if tlsCfg.Certificates[i], err = tls.LoadX509KeyPair(certFilePath, privFilePath); err != nil {
				beego.Error("Load cert file failure.", err)
				return err
			}
		}
	}

	server := &http.Server{Addr: addr, Handler: httpServeMux}
	server.SetKeepAlivesEnabled(true)

	beego.Info("start websocket wss listen: ", addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	if bWss {
		ln = tls.NewListener(ln, tlsCfg)
	}
	go func() {
		if err = server.Serve(ln); err != nil {
			beego.Error("server.Serve(\"", addr, "\") error(", err, ")")
			return
		}
	}()
	return nil
}

/*
* todo 阿里云slb的健康检测
 */
func (p *TcpGate) serveWebHealth(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(201)
}

/*
* todo websocket处理
 */
func (p *TcpGate) serveWebSocket(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		beego.Error("Websocket Upgrade error(", err, "), userAgent(", req.UserAgent(), ")")
		return
	}

	var (
		lAddr = ws.LocalAddr()
		rAddr = ws.RemoteAddr()
	)
	// nginx代理的客户端ip获取
	realIp := req.Header.Get("X-Real-Ip")

	if len(realIp) == 0 {
		// slb的客户端ip获取。X-Forwarded-For: 用户真实IP, 代理服务器1-IP， 代理服务器2-IP，...
		tmp := req.Header.Get("X-Forwarded-For")
		splitIdx := strings.Index(tmp, ",")
		if splitIdx > 0 {
			realIp = tmp[:splitIdx]
		} else {
			realIp = tmp
		}
	}

	session := link.NewSession(p.newWssCodec(atomic.AddUint32(&p.connSeed, 1), ws), p.config.SendChanSize)

	if conn, ok := session.Codec().(IGateConn); ok {
		conn.SetRealIp(realIp)
	}

	beego.Debug("start websocket serve", lAddr, "with", rAddr, ">>", realIp)

	p.handleSession(session, p.config.MaxConn)
}

// *********************************************
func (p *TcpGate) newTcpCodec(id uint32, conn net.Conn, bufferSize int) *TcpGateCodec {
	c := &TcpGateCodec{
		GateCodecBase: GateCodecBase{
			GateProtocol: &p.GateProtocol,
			id:           id,
			conn:         conn,
			reader:       bufio.NewReaderSize(conn, bufferSize),
		},
	}
	c.headBuf = c.headDat[:]
	return c
}

func (p *TcpGate) newTlsCodec(id uint32, conn net.Conn, bufferSize int) *TlsGateCodec {
	c := &TlsGateCodec{
		GateCodecBase: GateCodecBase{
			GateProtocol: &p.GateProtocol,
			id:           id,
			conn:         conn,
			reader:       bufio.NewReaderSize(conn, bufferSize),
		},
	}
	c.headBuf = c.headDat[:]
	return c
}

func (p *TcpGate) newWssCodec(id uint32, conn *websocket.Conn) *WsGateCodec {
	c := &WsGateCodec{
		GateProtocol: &p.GateProtocol,
		id:           id,
		conn:         conn,
	}
	c.headBuf = c.headDat[:]
	return c
}

// 处理客户端通讯，共用5个gorutine
// tcp 接收  --> 发送  grpc
// tcp 发送  <-- 接收  grpc
func (p *TcpGate) handleSession(session *link.Session, maxConn int) {
	co := session.Codec()
	conn, ok := co.(IGateConn)
	if !ok {
		return
	}
	id := conn.SocketID()

	clientip := conn.GetRealIp() // tool.Addr2IP(session.Codec().ClientAddr())
	beego.Info("OnConnect Session ip:", clientip, "id:", id)

	upChan := make(chan proto.Message, 100)
	downChan := make(chan proto.Message, 100)
	svrInfo := &GrpcSvrItem{}
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		close(upChan)
		close(downChan)

		if svrInfo.Ok && svrInfo.Conn != nil {
			svrInfo.Conn.Close()
		}

		beego.Info("Close Session ip:", clientip, "socketid:", id)
		if err := recover(); err != nil {
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(4)

	go p.handleTcpClient(session, conn, upChan, wg, ctx, cancel)

	go p.handleGrpcStream(svrInfo, upChan, downChan, wg, ctx, cancel)

	go p.handleDownGateMsg(session, downChan, wg, ctx, cancel)

	wg.Wait()
}

// 接受长连接的消息 gorutine
func (p *TcpGate) handleTcpClient(session *link.Session, tcpConn IGateConn, upChan chan proto.Message, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
	defer func() {
		wg.Done()
		cancel()
	}()

	id := tcpConn.SocketID()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if p.config.IdleTimeout > 0 {
			err := tcpConn.SetReadDeadline(time.Now().Add(p.config.IdleTimeout * time.Second))
			if err != nil {
				beego.Error("ReadDeadline time out idleTimeout:", p.config.IdleTimeout, err)
				return
			}
		}

		buf, err := session.Receive()
		if err != nil {
			beego.Error("Receive err Session id:", id, "err:", err)
			return
		}

		data := *(buf.(*[]byte))
		pkt := p.DecodePacket(data)
		p.Free(data)
		if err != nil {
			beego.Error("DecodePacket err session id:", id, "err:", err)
			return
		}

		packet, ok := pkt.(*GatePacket)
		if !ok {
			continue
		}
		gateMsg := &gate_msg.GatewayMsg{}
		if err := proto.Unmarshal(packet.Body, gateMsg); err != nil {
			return
		}
		// beego.Info("recvmsg:",string(packet.content),",sid:", id)
		upChan <- gateMsg
	}
}

// 接受grpc消息 gorutine
func (p *TcpGate) handleGrpcStream(svrInfo *GrpcSvrItem, upChan, downChan chan proto.Message, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
	defer func() {
		wg.Done()
		cancel()
	}()

	//var grpcClient gate_msg.CommonStreamSvrClient
	var grpcStream gate_msg.CommonStreamSvr_CommunicateClient
	var err error

	for {
		select {
		case <-ctx.Done():
			return
		case it := <-upChan:
			msg, ok := it.(*gate_msg.GatewayMsg)
			if !ok {
				return
			}
			if msg.Phase == gate_msg.GatewayMsgPhase_Communcate {
				if p.handleUpGrpcMsg(grpcStream, msg, downChan) == false {
					return
				}

			} else if msg.Phase == gate_msg.GatewayMsgPhase_Handshake {
				grpcStream, err = p.handleGrpcDial(msg, svrInfo)
				if err == nil {
					go p.handleDownGrpcMsg(grpcStream, downChan, svrInfo, wg, ctx, cancel)
					//
					ack := &gate_msg.GatewayMsg{
						Kind:    msg.Kind,
						Phase:   gate_msg.GatewayMsgPhase_Handshake,
						Message: "connect success",
					}
					downChan <- ack
				} else {
					ack := &gate_msg.GatewayMsg{
						Kind:    msg.Kind,
						Phase:   gate_msg.GatewayMsgPhase_Handshake,
						Message: "handshake fail",
						Code:    1,
					}
					downChan <- ack
					return
				}
			}
		}
	}

}

// 发起grpc服务连接
func (p *TcpGate) handleGrpcDial(msg *gate_msg.GatewayMsg, svrInfo *GrpcSvrItem) (gate_msg.CommonStreamSvr_CommunicateClient, error) {

	svrInfo.Name = msg.ServerName
	svrInfo.Kind = msg.Kind
	// 经过loadbalence 策略获取对应服务地址
	svrInfo.Addr = ""

	logs.Debug("svrInfo:", svrInfo)

	opts := grpc.WithInsecure()

	conn, err := grpc.Dial(svrInfo.Addr, opts)
	if err != nil {
		return nil, err
	}
	svrInfo.Conn = conn
	svrInfo.Ok = true

	grpcClient := gate_msg.NewCommonStreamSvrClient(conn)
	stream, err := grpcClient.Communicate(context.Background())
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// 发送消息给grpc服务
func (p *TcpGate) handleUpGrpcMsg(stream gate_msg.CommonStreamSvr_CommunicateClient, msg *gate_msg.GatewayMsg, down chan proto.Message) bool {
	streamMsg := &gate_msg.StreamMessage{}
	switch msg.GetKind() {
	case gate_msg.GatewayMsgKind_Json:
		if err := json.Unmarshal(msg.Body, streamMsg); err != nil {
			return false
		}
	case gate_msg.GatewayMsgKind_ProtoBuf:
		if err := proto.Unmarshal(msg.Body, streamMsg); err != nil {
			return false
		}
	}
	if stream != nil {
		stream.Send(streamMsg)
		return true
	} else {
		ack := &gate_msg.GatewayMsg{
			Kind:    gate_msg.GatewayMsgKind_ProtoBuf,
			Phase:   gate_msg.GatewayMsgPhase_Communcate,
			Message: "grpc stream error",
			Code:    1,
		}
		down <- ack
		return false
	}
}

// 接受grpc消息 gorutine
func (p *TcpGate) handleDownGrpcMsg(stream gate_msg.CommonStreamSvr_CommunicateClient, downChan chan proto.Message,
	svrInfo *GrpcSvrItem, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) error {
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			continue
		}
		body, err := proto.Marshal(in)
		if err != nil {
			continue
		}

		msg := &gate_msg.GatewayMsg{Kind: svrInfo.Kind, Phase: gate_msg.GatewayMsgPhase_Communcate, ServerName: svrInfo.Name, Body: body}
		downChan <- msg
	}
}

// 发送消息给创连接客户端 gorutine
func (p *TcpGate) handleDownGateMsg(session *link.Session, downChan chan proto.Message, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-downChan:
			body, err := proto.Marshal(msg)
			if err != nil {
				continue
			}
			pack := &GatePacket{
				Length: uint32(len(body)),
				Body:   body,
			}
			buffer := p.AllocPacket(pack)
			if err := session.Send(buffer); err != nil {

			}
		}
	}
}
