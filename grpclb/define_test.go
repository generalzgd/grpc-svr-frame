/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: policy.go
 * @time: 2019-07-18 14:01
 */

package grpclb

import "testing"

func TestValidatePolicyType(t *testing.T) {
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
			name:"TestValidatePolicyType",
			args:args{POLICY_LOW_FIRST},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePolicyType(tt.args.in); got != tt.want {
				t.Errorf("ValidatePolicyStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
