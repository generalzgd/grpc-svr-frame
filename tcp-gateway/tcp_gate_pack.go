/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: tcp_gate_pack.go
 * @time: 2019-07-17 23:08
 */

package tcp_gateway

import (
	"encoding/binary"

	"github.com/funny/slab"
)

type GatePacket struct {
	Length uint32
	Ver    uint16
	Opt    uint16
	Body   []byte
}

type PackEncrypt struct {
	HeadSize      int
	MaxPacketSize int
	Pool          slab.Pool
}

// 重置内存块
func (p *PackEncrypt) MemSet(mem []byte, v byte) {
	l := len(mem)
	for i := 0; i < l; i += 1 {
		mem[i] = v
	}
}

func (p *PackEncrypt) Init(headSize, maxPacketSize int, pool slab.Pool) {
	p.HeadSize = headSize
	p.MaxPacketSize = maxPacketSize
	p.Pool = pool
}

func (p *PackEncrypt) Alloc(size int) []byte {
	if p.Pool != nil {
		return p.Pool.Alloc(size)
	}
	return make([]byte, size)
}

func (p *PackEncrypt) Free(msg []byte) {
	if p.Pool != nil {
		p.Pool.Free(msg)
	}
}

func (p *PackEncrypt) AllocPacket(pack interface{}) []byte {
	return []byte{}
}
func (p *PackEncrypt) DecodePacket(data []byte) interface{} {
	return struct{}{}
}

// 网关包编解码器
type GatePackEncrypt struct {
	PackEncrypt
}

// 在传入之后，会重新计算body，以及body的长度
func (p *GatePackEncrypt) AllocPacket(pkt interface{}) []byte {
	gatePack, ok := pkt.(*GatePacket)
	if !ok {
		return []byte{}
	}
	buffer := p.Alloc(p.HeadSize + len(gatePack.Body))

	binary.LittleEndian.PutUint32(buffer[0:], uint32(len(gatePack.Body)))
	binary.LittleEndian.PutUint16(buffer[4:], gatePack.Ver)
	binary.LittleEndian.PutUint16(buffer[6:], gatePack.Opt)

	copy(buffer[p.HeadSize:], gatePack.Body)

	return buffer
}

// decodePacket decodes gateway message
func (p *GatePackEncrypt) DecodePacket(data []byte) interface{} {
	pack := &GatePacket{}
	pack.Length = binary.LittleEndian.Uint32(data[0:])
	pack.Body = make([]byte, pack.Length)
	copy(pack.Body, data[p.HeadSize:])

	return pack
}




type GateProtocol struct {
	GatePackEncrypt
}