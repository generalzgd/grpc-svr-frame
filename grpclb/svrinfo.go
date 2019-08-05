/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: svrinfo.go
 * @time: 2019-07-18 12:32
 */

package grpclb

import (
	"sync"

	"github.com/generalzgd/grpc-svr-frame/config"
	"github.com/generalzgd/grpc-svr-frame/tcp-gateway/proto"
)

// 单个grpc服务实例结构
type GrpcSvrInfo struct {
	Name      string                  //
	Addr      string                  // ip:port
	Kind      gate_msg.GatewayMsgKind //
	Insecure  bool                    // 是否启用tls
	CertFiles []config.CertFile       //
}

// 同类型的grpc 集合
type GrpcSvrGroup struct {
	lock    sync.Mutex
	SvrData map[string]*GrpcSvrInfo
}
