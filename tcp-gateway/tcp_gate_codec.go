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
	"bytes"
	"encoding/binary"
	"io"
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

type GateCodecBase struct {
	*GateProtocol
	id      uint32
	conn    net.Conn
	reader  *bufio.Reader
	headBuf []byte
	headDat [TCP_GATE_PACKAGE_HEAD_LENGTH]byte
}

func (p *GateCodecBase) Receive() (interface{}, error) {
	//panic("implement me")
	p.MemSet(p.headBuf, 0)
	if _, err := io.ReadFull(p.reader, p.headBuf);err != nil {
		return nil, err
	}

	length := int(binary.LittleEndian.Uint32(p.headBuf[0:]))
	if length > p.MaxPacketSize {
		return nil, bytes.ErrTooLarge
	}

	buffer := p.Alloc(TCP_GATE_PACKAGE_HEAD_LENGTH + length)
	copy(buffer, p.headBuf)
	if _, err := io.ReadFull(p.reader, buffer[TCP_GATE_PACKAGE_HEAD_LENGTH:]); err != nil {
		p.Free(buffer)
		return nil, err
	}
	return &buffer, nil
}

type TcpGateCodec struct {

	GateCodecBase
	realIp  string
}

func (p *TcpGateCodec) ClearSendChan(sendChan <-chan interface{}) {
	//panic("implement me")
	for msg := range sendChan {
		p.Free(*(msg.(*[]byte)))
	}
}

func (p *TcpGateCodec) Receive() (interface{}, error) {
	//panic("implement me")
	return p.GateCodecBase.Receive()
}

func (p *TcpGateCodec) Send(msg interface{}) error {
	//panic("implement me")
	buffer := *(msg.(*[]byte))
	_, err := p.conn.Write(buffer)
	p.Free(buffer)
	return err
}

func (p *TcpGateCodec) Close() error {
	//panic("implement me")
	return p.conn.Close()
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
	GateCodecBase
	realIp  string
}

func (p *TlsGateCodec) ClearSendChan(sendChan <-chan interface{}) {
	//panic("implement me")
	for msg := range sendChan {
		p.Free(*(msg.(*[]byte)))
	}
}

func (p *TlsGateCodec) Receive() (interface{}, error) {
	//panic("implement me")
	return p.GateCodecBase.Receive()
}

func (p *TlsGateCodec) Send(msg interface{}) error {
	//panic("implement me")
	buffer := *(msg.(*[]byte))
	_, err := p.conn.Write(buffer)
	p.Free(buffer)
	return err
}

func (p *TlsGateCodec) Close() error {
	//panic("implement me")
	return p.conn.Close()
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

func (p *WsGateCodec) SocketID() uint32 {
	//panic("implement me")
	return p.id
}

func (p * WsGateCodec) ClientAddr() string {
	//panic("implement me")
	return p.conn.RemoteAddr().String()
}

func (p * WsGateCodec) SetReadDeadline(t time.Time) error {
	//panic("implement me")
	return p.conn.SetReadDeadline(t)
}

func (p * WsGateCodec) SetRealIp(v string) {
	//panic("implement me")
	p.realIp = v
}

func (p * WsGateCodec) GetRealIp() string {
	//panic("implement me")
	return p.realIp
}

func (p * WsGateCodec) ClearSendChan(sendChan <-chan interface{}) {
	//panic("implement me")
	for msg := range sendChan {
		p.Free(*(msg.(*[]byte)))
	}
}

func (p * WsGateCodec) Receive() (interface{}, error) {
	//panic("implement me")
	_, r, err := p.conn.NextReader()
	if err != nil {
		return nil, err
	}

	if _, err := io.ReadFull(r, p.headBuf);err != nil {
		return nil, err
	}

	length := int(binary.LittleEndian.Uint32(p.headBuf[0:]))
	if length > p.MaxPacketSize {
		return nil, bytes.ErrTooLarge
	}

	buffer := p.Alloc(TCP_GATE_PACKAGE_HEAD_LENGTH + length)
	copy(buffer, p.headBuf)
	if _, err := io.ReadFull(r, buffer[TCP_GATE_PACKAGE_HEAD_LENGTH:]); err != nil {
		p.Free(buffer)
		return nil, err
	}
	return &buffer, nil
}

func (p * WsGateCodec) Send(msg interface{}) error {
	//panic("implement me")
	buffer := *(msg.(*[]byte))
	err := p.conn.WriteMessage(2, buffer)
	p.Free(buffer)
	return err
}

func (p * WsGateCodec) Close() error {
	//panic("implement me")
	return p.conn.Close()
}

