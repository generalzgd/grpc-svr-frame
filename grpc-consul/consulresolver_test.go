/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: balance.go
 * @time: 2019/9/11 13:47
 */
package grpc_consul

import (
	"testing"
	"time"
)

func TestInitRegister(t *testing.T) {
	type args struct {
		consulAddr string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "TestInitRegister",
			args: args{
				consulAddr: "http://127.0.0.1:8500",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitRegister(tt.args.consulAddr)
		})
	}

	time.Sleep(time.Hour)
}

func Test_testCallRpc(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name:"Test_testCallRpc_t1",
		},
	}
	for _, tt := range tests {
		// for i:=0;i<2;i++{
			t.Run(tt.name, func(t *testing.T) {
				// testCallRpc()
			})
		// }

	}
}
