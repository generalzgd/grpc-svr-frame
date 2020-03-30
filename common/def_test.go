/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: def.go
 * @time: 2019-07-18 15:30
 */

package common

import "testing"

func TestValidateGatewayType(t *testing.T) {
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
			name: "t1",
			args: args{
				in: GW_TYPE_HTTP,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateGatewayType(tt.args.in); got != tt.want {
				t.Errorf("ValidateGatewayType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateLbPolicyType(t *testing.T) {
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
			name: "t1",
			args: args{LB_POLICY_LOW_FIRST},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateLbPolicyType(tt.args.in); got != tt.want {
				t.Errorf("ValidateLbPolicyType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateStasticsType(t *testing.T) {
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
			name: "t1",
			args: args{STI_TYPE_TPS},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateStasticsType(tt.args.in); got != tt.want {
				t.Errorf("ValidateStasticsType() = %v, want %v", got, tt.want)
			}
		})
	}
}
