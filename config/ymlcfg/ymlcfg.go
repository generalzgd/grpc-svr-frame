/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: cfg.go
 * @time: 2019-08-19 14:22
 */

package ymlcfg

import (
	`time`
)

// grpc链接池配置
type ConnPool struct {
	InitNum     int           `yaml:"init"`        // 初始数量
	CapNum      int           `yaml:"cap"`         // 最对连接数
	IdleTimeout time.Duration `yaml:"idletimeout"` // 空闲超时
	LifeDur     time.Duration `yaml:"livedur"`     // 最大生命周期
}

func (p *ConnPool) GetInitNum() int {
	if p.InitNum < 1 {
		return 1
	}
	return p.InitNum
}

func (p *ConnPool) GetCapNum() int {
	if p.CapNum < 1 {
		return 10
	}
	return p.CapNum
}

type MemPoolConfig struct {
	Type     string `yaml:"type"`
	Factor   int    `yaml:"factor"`
	MinChunk int    `yaml:"minChunk"`
	MaxChunk int    `yaml:"maxChunk"`
	PageSize int    `yaml:"pageSize"`
}

type ConsulConfig struct {
	Address    string `yaml:"address"`
	Token      string `yaml:"token"`
	HealthType string `yaml:"healthtype"` // 1 tcp 2 http
	HealthPort int    `yaml:"healthport"` //
}

type EndpointConfig struct {
	Name      string     `yaml:"name"`
	Address   string     `yaml:"address"` // 客户端用该字段, dns名:port/ 服务名[:port]/ ip:port
	Port      int        `yaml:"port"`    // 服务端用该字段
	Secure    bool       `yaml:"secure"`
	CertFiles []CertFile `yaml:"certfiles"`
	Pool      ConnPool   `yaml:"pool"` // 链接池配置
}

type CertFile struct {
	Cert       string `yaml:"cert"` // 证书文件，pem格式
	Priv       string `yaml:"priv"` // 密钥
	RootCAFile string `yaml:"ca"`   // ca证书
}
