/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: tcp_gate_config.go
 * @time: 2019-07-17 23:08
 */

package tcp_gateway

import "time"

type CertFile struct {
	Cert string `toml:"cert"` // 证书文件，pem格式
	Priv string `toml:"priv"` // 密钥
}

type GateConnConfig struct {
	Insecure        bool       // false: tcp/ws  true: tls/wss, 版本号默认1.1
	UseWs bool // false: tcp true: ws
	CertFiles       []CertFile // 证书文件，pem格式
	BufferSize      int
	MaxConn         int
	IdleTimeout     time.Duration
	SendChanSize    int    //
	ReceiveChanSize int    //
	Port            uint32 // 侦听端口
}

