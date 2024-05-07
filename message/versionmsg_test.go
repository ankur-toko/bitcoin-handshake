package message

import (
	"net"
	"testing"
	"time"
)

func TestVersionMsg_ToBytes(t *testing.T) {
	type fields struct {
		Version     uint32
		Timestamp   uint64
		Services    uint64
		RecAddr     NetAddress
		SendAddr    NetAddress
		Nounce      uint64
		UserAgent   string
		StartHeight uint32
		Relay       bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"incorrect version",
			fields{
				12,
				uint64(time.Now().Unix()),
				0,
				NetAddress{},
				NetAddress{},
				0,
				"dummy",
				0,
				false,
			},
			nil,
			true,
		}, {
			"correct version message",
			fields{
				70017,
				uint64(time.Now().Unix()),
				0,
				NetAddress{time.Now(), 0, net.IP("10.0.0.0"), 8924},
				NetAddress{},
				0,
				"dummy",
				0,
				false,
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VersionMsg{
				Version:     tt.fields.Version,
				Timestamp:   tt.fields.Timestamp,
				Services:    tt.fields.Services,
				RecAddr:     tt.fields.RecAddr,
				SendAddr:    tt.fields.SendAddr,
				Nounce:      tt.fields.Nounce,
				UserAgent:   tt.fields.UserAgent,
				StartHeight: tt.fields.StartHeight,
				Relay:       tt.fields.Relay,
			}
			_, err := v.ToBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionMsg.ToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("VersionMsg.ToBytes() = %v, want %v", got, tt.want)
			// }
		})
	}
}
