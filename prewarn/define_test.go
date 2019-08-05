/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: define.go
 * @time: 2019-07-30 15:48
 */

package prewarn

import "testing"

func TestValidateWarnType(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name:"TestValidateWarnType",
			args:args{WARN_EMAIL},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateWarnType(tt.args.in); got != tt.want {
				t.Errorf("ValidateWarnType() = %v, want %v", got, tt.want)
			}
		})
	}
}
