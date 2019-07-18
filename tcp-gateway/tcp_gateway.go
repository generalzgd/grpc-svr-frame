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
	"crypto/tls"
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
	"github.com/gorilla/websocket"
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


type TcpGate struct {
	GateProtocol
	config   GateConnConfig
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

func (p *TcpGate) Initialize() {
	makePool := func(t string, minChunk, maxChunk, factor, pageSize int) slab.Pool {
		switch t {
		case "sync":
			return slab.NewSyncPool(minChunk, maxChunk, factor)
		case "atom":
			return slab.NewAtomPool(minChunk, maxChunk, factor, pageSize)
		case "chan":
			return slab.NewChanPool(minChunk, maxChunk, factor, pageSize)
		default:
			beego.Error(`unsupported memory pool type, must be "sync", "atom" or "chan"`)
		}
		return nil
	}

	pool := makePool("atom", 64, 65536, 2, 1048576)
	p.PackEncrypt.Init(TCP_GATE_PACKAGE_HEAD_LENGTH, CONTENT_LENGTH_LIMIT, pool)
	p.exit = make(chan struct{})
}

func (p *TcpGate) Destroy() {
	p.closeOnce.Do(func() {
		close(p.exit)
		p.closeFlag = true
	})
}

func (p *TcpGate) Run(workerNum int) {
}

func (p *TcpGate) Serve(cfg GateConnConfig) error {
	if cfg.UseWs {
		return p.ServeWs(cfg)
	} else {
		return p.ServeTcp(cfg)
	}
}

func (p *TcpGate) ServeTcp(cfg GateConnConfig) error {
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
				p.handleSession(session, cfg.MaxConn, cfg.IdleTimeout)
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
				p.handleSession(session, cfg.MaxConn, cfg.IdleTimeout)
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

func (p *TcpGate) ServeWs(cfg GateConnConfig) error {
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

	p.handleSession(session, p.config.MaxConn, p.config.IdleTimeout)
}

// *********************************************
func (p *TcpGate) newTcpCodec(id uint32, conn net.Conn, bufferSize int) *TcpGateCodec {
	c := &TcpGateCodec{
		GateProtocol: &p.GateProtocol,
		id:           id,
		conn:         conn,
		reader:       bufio.NewReaderSize(conn, bufferSize),
	}
	c.headBuf = c.headDat[:]
	return c
}

func (p *TcpGate) newTlsCodec(id uint32, conn net.Conn, bufferSize int) *TlsGateCodec {
	c := &TlsGateCodec{
		GateProtocol: &p.GateProtocol,
		id:           id,
		conn:         conn,
		reader:       bufio.NewReaderSize(conn, bufferSize),
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

func (p *TcpGate) handleSession(session *link.Session, maxConn int, idleTimeout time.Duration) {
	co := session.Codec()
	conn, ok := co.(IGateConn)
	if !ok {
		return
	}
	id := conn.SocketID()

	clientip := conn.GetRealIp() // tool.Addr2IP(session.Codec().ClientAddr())
	beego.Info("OnConnect Session ip:", clientip, "id:", id)

	defer func() {
		beego.Info("Close Session ip:", clientip, "socketid:", id)
		if err := recover(); err != nil {
		}
	}()

	for {
		if idleTimeout > 0 {
			err := conn.SetReadDeadline(time.Now().Add(idleTimeout * time.Second))
			if err != nil {
				beego.Error("ReadDeadline time out idleTimeout:", idleTimeout, err)
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

		// beego.Info("recvmsg:",string(packet.content),",sid:", id)

	}
}
