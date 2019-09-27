/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: analyse.go
 * @time: 2019/9/25 14:19
 */
package analyse

import (
	"reflect"
	"testing"

	`github.com/generalzgd/grpc-svr-frame/monitor`
)

func TestNewAnalyse(t *testing.T) {
	type args struct {
		threshold   uint
		analyseNum  uint
		analyseType string
		fieldName   string
	}
	tests := []struct {
		name string
		args args
		want *Analyse
	}{
		// TODO: Add test cases.
		{
			name: "TestNewAnalyse",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyse(tt.args.threshold, tt.args.analyseNum, tt.args.analyseType, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyse_analyseData(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		args   args
	}{
		// TODO: Add test cases.
		{
			name:"TestAnalyse_analyseData",
			args:args{
				data:`{"foo":{“boo”:568}, "a":true, "b":0.235}`,
			},
		},
	}
	p := NewAnalyse(1, 1, monitor.ANALYSE_CNT, "boo")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.analyseData(tt.args.data)
		})
	}
}
