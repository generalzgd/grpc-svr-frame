/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: watch.go
 * @time: 2019/8/22 8:52
 */
package monitor

import (
	"testing"
)

func TestNewRecord(t *testing.T) {
	type args struct {
		typ  int
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name:"TestNewRecord",
			args:args{
				typ:Stat_Mem,
				args: []interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewRecord(tt.args.typ, tt.args.args...)
		})
	}
}
