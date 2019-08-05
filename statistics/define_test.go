/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: define.go
 * @time: 2019-07-30 15:42
 */

package statistics

import "testing"

func TestValidateStatisticType(t *testing.T) {
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
			name:"TestValidateStatisticType",
			args:args{STS_CNT},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateStatisticType(tt.args.in); got != tt.want {
				t.Errorf("ValidateStatisticType() = %v, want %v", got, tt.want)
			}
		})
	}
}
