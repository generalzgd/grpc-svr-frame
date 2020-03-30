/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: config.go
 * @time: 2019-07-18 12:38
 */

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type CertFile struct {
	Cert string `yaml:"cert"` // 证书文件，pem格式
	Priv string `yaml:"priv"` // 密钥
}

type PolicyPercentConfig struct {
	Endpoint string  `yaml:"endpoint"`
	Percent  float64 `yaml:"percent"` // 总和不超过100%
}

// 网关连接配置
type GateConnConfig struct {
	Insecure        bool          `yaml:"insecure"`  // false: http/tcp/ws  true: https/tls/wss, 版本号默认1.1
	CertFiles       []CertFile    `yaml:"certfiles"` // 证书文件，pem格式
	BufferSize      int           `yaml:"buffersize"`
	MaxConn         int           `yaml:"maxconn"`
	IdleTimeout     time.Duration `yaml:"idletimeout"`
	SendChanSize    int           `yaml:"sendchansize"` //
	ReceiveChanSize int           `yaml:"recvchansize"` //
	Port            uint32        `yaml:"port"`         // 侦听端口
}

// 后端grpc服务配置
type EndpointConfig struct {
	Insecure  bool       `yaml:"insecure"`
	CertFiles []CertFile `yaml:"certfiles"`
	Endpoint  string     `yaml:"endpoint"` // 目的服务地址 IP:Port
}

// 转发策略
type PolicyConfig struct {
	Type       string                `yaml:"type"`    // 策略类型
	PercetConf []PolicyPercentConfig `yaml:"percent"` // 指定比例策略配置
}

type StatisticsConfig struct {
	Type  string `yaml:"type"`  // 统计类型
	Field string `yaml:"field"` // 指定字段 sum/avg -> number, count -> string
}

// 单点预警 tps
type EndpointWarnConfig struct {
	Endpoint  string  `yaml:"endpoint"`
	Threshold float64 `yaml:"threshold"` // tps warning
	Method    string  `yaml:"method"`    // sms/email
}

// grpc服务预警 tps
type SvrWarnConfig struct {
	Endpoint  string  `yaml:"endpoint"`
	Threshold float64 `yaml:"threshold"` // tps warning
	Method    string  `yaml:"method"`    // sms/email
}

// 网关负载预警 tps
type NodeWarnConfig struct {
	Threshold float64 `yaml:"threshold"` // tps warning
	Method    string  `yaml:"method"`    // sms/email
}

type EarlyWarnConfig struct {
	EndpointWarn []EndpointWarnConfig `yaml:"endpointwarn"` // 目标服务点
	SvrWarn      []SvrWarnConfig      `yaml:"svrwarn"`      // 目标服务
}

// 单项网关连接配置
type GwItemConfig struct {
	Type         string           `yaml:"type"` // 网关协议类型  GW_TYPE_xxx
	GwConf       GateConnConfig   `yaml:"gwconf"`
	EndpointConf EndpointConfig   `yaml:"endpointconf"`
	PolicyConf   PolicyConfig     `yaml:"policyconf"`
	StatiConf    StatisticsConfig `yaml:"sticonf"`
	EarWarnConf  EarlyWarnConfig  `yaml:"warnconf"`
}

// 网关整体配置
type GatewayConfig struct {
	Name     string         `yaml:"name"`
	Ver      string         `yaml:"ver"`
	NodeWarn NodeWarnConfig `yaml:"nodewarn"`
	List     []GwItemConfig `yaml:"list"`
}

func (p *GatewayConfig) Load(path string) error {
	if len(path) < 1 {
		path = filepath.Join(filepath.Dir(os.Args[0]), "config", fmt.Sprintf("config_%s.yaml", "dev"))
	}
	// file := filepath.Join(path, "config", fmt.Sprintf("config_%s.yaml", "dev"))

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	bts, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bts, p); err != nil {
		return err
	}

	return nil
}
