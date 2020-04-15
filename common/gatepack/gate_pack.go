/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: client_package.go
 * @time: 2019-08-04 19:05
 */

package gatepack

import (
	"bytes"
	"encoding/binary"
	`fmt`
)

const (
	ProtocolVer = 100 // 1.0.0

	PackCodecProto = 0
	PackCodecJson  = 1

	PackHeadSize = 12
	PackMaxSize  = 1024 * 32
)

type GateClientPackHead struct {
	Length uint16 // body的长度，65535/1024 ~ 63k
	Seq    uint16 // 序列号
	CmdId  uint16 // 协议id，可以映射到对应的service:method（兼容字段，后期考虑，把房间聊天网关迁移过来）
	Ver    uint16 // 协议更新版本号 1.0.1 => 1*100 + 0*10 + 1 => 101
	Codec  uint16 // 0:proto  1:json
	Opt    uint16 // 备用字段
}

// 网关包, 小端
type GateClientPack struct {
	GateClientPackHead
	Body []byte // protobuf or json
}

func (p *GateClientPack) Serialize() []byte {
	out, _ := p.SerializeWithBuf(nil)
	return out
}

func (p *GateClientPack) SerializeWithBuf(out []byte) ([]byte, error) {
	// val := int(unsafe.Sizeof(p.GateClientPackHead))
	if cap(out) < 1 {
		out = make([]byte, 0, PackHeadSize+p.Length)
	}

	buf := bytes.NewBuffer(out)
	if err := binary.Write(buf, binary.LittleEndian, p.GateClientPackHead); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, p.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *GateClientPack) Deserialize(in []byte) error {
	r := bytes.NewReader(in)
	if err := binary.Read(r, binary.LittleEndian, &p.GateClientPackHead); err != nil {
		return err
	}
	p.Body = make([]byte, p.Length)
	copy(p.Body, in[PackHeadSize:])
	return nil
}

func (p *GateClientPack) String()string {
	return fmt.Sprintf("head:%v,Body:%s", p.GateClientPackHead, string(p.Body))
}