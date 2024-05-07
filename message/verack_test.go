package message

import (
	"reflect"
	"testing"
)

func TestVerAck_ToBytes(t *testing.T) {
	v := VerAck{}

	tests := []struct {
		name    string
		v       VerAck
		want    []byte
		wantErr bool
	}{
		{
			"basic",
			v,
			[]byte{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VerAck{}
			got, err := v.ToBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("VerAck.ToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerAck.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
