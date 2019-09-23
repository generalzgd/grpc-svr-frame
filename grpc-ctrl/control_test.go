/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: control.go
 * @time: 2019/8/12 15:07
 */
package grpc_ctrl

import (
	"reflect"
	"testing"

	`google.golang.org/grpc`
)

func TestMakeGrpcController(t *testing.T) {
	tests := []struct {
		name string
		want GrpcController
	}{
		// TODO: Add test cases.
		{
			name:"make controller",
			want:GrpcController{grpcConnMap: map[string]*grpc.ClientConn{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeGrpcController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeGrpcController() = %v, want %v", got, tt.want)
			}
		})
	}
}
