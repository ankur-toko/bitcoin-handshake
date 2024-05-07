package checksum

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestDoubleSha256Hash(t *testing.T) {
	type args struct {
		data []byte
	}

	decode, _ := hex.DecodeString("5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456")
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "empty",
			args: args{
				data: []byte{},
			},
			want: decode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleSha256Hash(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				// v := fmt.Sprintf("%x", got)
				t.Errorf("DoubleSha256Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
