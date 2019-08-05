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

import "testing"

func TestGatewayConfig_Load(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:"t1",
			args:args{
				path:"/Users/zgd/WorkSpace/GoSpaceMod/grpc-svr-frame",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GatewayConfig{}
			if err := p.Load(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("GatewayConfig.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
