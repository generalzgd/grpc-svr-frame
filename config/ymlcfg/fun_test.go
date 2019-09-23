/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: fun.go
 * @time: 2019-08-03 22:09
 */

package ymlcfg

import "testing"

func TestLoadYaml(t *testing.T) {
	type args struct {
		path string
		tar  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestLoadYaml",
			args: args{
				path: "",
				tar:  nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadYaml(tt.args.path, tt.args.tar); (err != nil) != tt.wantErr {
				t.Errorf("LoadYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
