/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: tcp_conn.go
 * @time: 2019-07-17 23:07
 */

package tcp_gateway

import (
	"bufio"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

type IGateConn interface {
	SocketID() uint32
	ClientAddr() string
	SetReadDeadline(t time.Time) error
	SetRealIp(v string)
	GetRealIp() string
}

type TcpGateCodec struct {
	*GateProtocol
	id      uint32
	conn    net.Conn
	reader  *bufio.Reader
	headBuf []byte
	headDat [TCP_GATE_PACKAGE_HEAD_LENGTH]byte
	realIp  string
}

func (p *TcpGateCodec) ClearSendChan(<-chan interface{}) {
	panic("implement me")
}

func (p *TcpGateCodec) Receive() (interface{}, error) {
	panic("implement me")
}

func (p *TcpGateCodec) Send(interface{}) error {
	panic("implement me")
}

func (p *TcpGateCodec) Close() error {
	panic("implement me")
}

func (p *TcpGateCodec) SocketID() uint32 {
	return p.id
}

func (p *TcpGateCodec) ClientAddr() string {
	return p.conn.RemoteAddr().String()
}

func (p *TcpGateCodec) SetReadDeadline(t time.Time) error {
	return p.conn.SetReadDeadline(t)
}

func (p *TcpGateCodec) SetRealIp(v string) {
	p.realIp = v
}

func (p *TcpGateCodec) GetRealIp() string {
	return p.realIp
}

// //////////////////////////
type TlsGateCodec struct {
	*GateProtocol
	id      uint32
	conn    net.Conn
	reader  *bufio.Reader
	headBuf []byte
	headDat [TCP_GATE_PACKAGE_HEAD_LENGTH]byte
	realIp  string
}

func (TlsGateCodec) ClearSendChan(<-chan interface{}) {
	panic("implement me")
}

func (TlsGateCodec) Receive() (interface{}, error) {
	panic("implement me")
}

func (TlsGateCodec) Send(interface{}) error {
	panic("implement me")
}

func (TlsGateCodec) Close() error {
	panic("implement me")
}

func (p *TlsGateCodec) SocketID() uint32 {
	return p.id
}

func (p *TlsGateCodec) ClientAddr() string {
	return p.conn.RemoteAddr().String()
}

func (p *TlsGateCodec) SetReadDeadline(t time.Time) error {
	return p.conn.SetReadDeadline(t)
}

func (p *TlsGateCodec) SetRealIp(v string) {
	p.realIp = v
}

func (p *TlsGateCodec) GetRealIp() string {
	return p.realIp
}

// ////////
type WsGateCodec struct {
	*GateProtocol
	id      uint32
	conn    *websocket.Conn
	headBuf []byte
	headDat [TCP_GATE_PACKAGE_HEAD_LENGTH]byte
	realIp  string
}

func ( WsGateCodec) SocketID() uint32 {
	panic("implement me")
}

func ( WsGateCodec) ClientAddr() string {
	panic("implement me")
}

func ( WsGateCodec) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func ( WsGateCodec) SetRealIp(v string) {
	panic("implement me")
}

func ( WsGateCodec) GetRealIp() string {
	panic("implement me")
}

func ( WsGateCodec) ClearSendChan(<-chan interface{}) {
	panic("implement me")
}

func ( WsGateCodec) Receive() (interface{}, error) {
	panic("implement me")
}

func ( WsGateCodec) Send(interface{}) error {
	panic("implement me")
}

func ( WsGateCodec) Close() error {
	panic("implement me")
}

