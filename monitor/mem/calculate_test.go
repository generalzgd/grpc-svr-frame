/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: calculate.go
 * @time: 2019/8/22 11:04
 */
package mem

import (
	"testing"
)

func Test_formateMem(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name:"Test_formateMem_1",
			args:args{
				num:5*1024*1024,
			},
			want:"5M",
		},
		{
			name:"Test_formateMem_2",
			args:args{
				num:5*1024*1024+23,
			},
			want:"5M23B",
		},
		{
			name:"Test_formateMem_3",
			args:args{
				num:5*1024*1024+1024+23,
			},
			want:"5M1K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formateMem(tt.args.num); got != tt.want {
				t.Errorf("formateMem() = %v, want %v", got, tt.want)
			}
		})
	}
}
