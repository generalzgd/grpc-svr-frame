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
}

type CertFile struct {
	Cert       string `yaml:"cert"` // 证书文件，pem格式
	Priv       string `yaml:"priv"` // 密钥
	RootCAFile string `yaml:"ca"`   // ca证书
}
