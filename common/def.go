/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: def.go
 * @time: 2019-07-18 15:30
 */

package common

import (
	"strings"
	`time`

	"github.com/toolkits/slice"
)

const (
	LbPolicyRandom   = "random"   // 随机 http默认方式
	LbPolicyHold     = "hold"     // 会话保持  长链接默认方式
	LbPolicyPercent  = "percent"  // 指定比例
	LbPolicyLowFirst = "lowfirst" // 低负载优先
)

const (
	GwTypeHttp = "http" // http1.1 http2(grpc直接支持，gateway则不再实现转发，loadbanch需要实现)
	GwTypeTcp  = "tcp"  // 长短连接
	GwTypeWs   = "ws"   // 长短连接
)

const (
	StiTypeSum   = "sum" // 统计类型枚举
	StiTypeAvg   = "avg"
	StiTypeCount = "count"
	StiTypeTps   = "tps"
)

var (
	GwTypeList   = []string{GwTypeHttp, GwTypeTcp, GwTypeWs}
	LbPolicyList = []string{LbPolicyRandom, LbPolicyHold, LbPolicyPercent, LbPolicyLowFirst}
	StiTypeList  = []string{StiTypeSum, StiTypeAvg, StiTypeCount, StiTypeTps}
)

func ValidateGatewayType(in string) bool {
	in = strings.ToLower(in)
	return slice.ContainsString(GwTypeList, in)
}

func ValidateLbPolicyType(in string) bool {
	in = strings.ToLower(in)
	return slice.ContainsString(LbPolicyList, in)
}

func ValidateStasticsType(in string) bool {
	in = strings.ToLower(in)
	return slice.ContainsString(StiTypeList, in)
}


// 客户端连接的用户信息
type ClientConnInfo struct {
	SocketId  uint32 `json:"socket_id,omitempty"`
	ClientIp  string `json:"client_ip,omitempty"`
	GateIp    string `json:"gate_ip,omitempty"`
	LoginTime int64  `json:"login_time,omitempty"` // 上次登录时间
	Expire    int64  `json:"-"`                    // 过期时间 s，logintime + expire < now, expire==0 不过期
	Guid string `json:"-"`
	Uid uint32 `json:"uid"`
	Platform uint32 `json:"platform"`
	State bool `json:"state"`
	Nickname string `json:"nickname"`
}

// 是否过期
func (p *ClientConnInfo) Expired() bool {
	if p.Expire > 0 {
		if p.LoginTime+p.Expire < time.Now().Unix() {
			return true
		}
	}
	return false
}