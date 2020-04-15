/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: client_package.go
 * @time: 2019-08-04 19:05
 */

package gatepack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGateClientPack_Serialize(t *testing.T) {
	fmt.Println(ProtocolVer, PackCodecProto, PackCodecJson, PackMaxSize)
	type fields struct {
		Length uint16
		Seq    uint16
		Cmdid  uint16
		Ver    uint16
		Codec  uint16
		Opt    uint16
		Body   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Serialize",
			fields: fields{
				Length: 9,
				Seq:    1,
				Cmdid:  266,
				Ver:    1,
				Codec:  0,
				Opt:    0,
				Body:   []byte{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GateClientPack{
				GateClientPackHead: GateClientPackHead{
					Length: tt.fields.Length,
					Seq:    tt.fields.Seq,
					Cmdid:  tt.fields.Cmdid,
					Ver:    tt.fields.Ver,
					Codec:  tt.fields.Codec,
					Opt:    tt.fields.Opt,
				},
				Body: tt.fields.Body,
			}
			got := p.Serialize()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GateClientPack.Serise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGateClientPack_Unserialize(t *testing.T) {
	type fields struct {
		GateClientPackHead GateClientPackHead
		Body               []byte
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestGateClientPack_Unserialize",
			fields: fields{},
			args: args{
				in: []byte{9, 0, 1, 0, 10, 1, 1, 0, 0, 0, 0, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GateClientPack{
				GateClientPackHead: tt.fields.GateClientPackHead,
				Body:               tt.fields.Body,
			}
			if err := p.Unserialize(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("GateClientPack.Unserialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
