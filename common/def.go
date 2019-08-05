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

	"github.com/toolkits/slice"
)

const (
	LB_POLICY_RANDOM    = "random"   // 随机 http默认方式
	LB_POLICY_HOLD      = "hold"     // 会话保持  长链接默认方式
	LB_POLICY_PERCENT   = "percent"  // 指定比例
	LB_POLICY_LOW_FIRST = "lowfirst" // 低负载优先
)

const (
	GW_TYPE_HTTP = "http" // http1.1 http2(grpc直接支持，gateway则不再实现转发，loadbanch需要实现)
	GW_TYPE_TCP  = "tcp"  // 长短连接
	GW_TYPE_WS   = "ws"   // 长短连接
)

const (
	STI_TYPE_SUM   = "sum" // 统计类型枚举
	STI_TYPE_AVG   = "avg"
	STI_TYPE_COUNT = "count"
	STI_TYPE_TPS   = "tps"
)

var (
	GwTypeList   = []string{GW_TYPE_HTTP, GW_TYPE_TCP, GW_TYPE_WS}
	LbPolicyList = []string{LB_POLICY_RANDOM, LB_POLICY_HOLD, LB_POLICY_PERCENT, LB_POLICY_LOW_FIRST}
	StiTypeList  = []string{STI_TYPE_SUM, STI_TYPE_AVG, STI_TYPE_COUNT, STI_TYPE_TPS}
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
