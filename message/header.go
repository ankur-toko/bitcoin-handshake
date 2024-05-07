package message

import (
	"encoding/binary"

	"github.com/ankur-toko/bitcoin-handshake/checksum"
)

type Header struct {
	MagicWord  uint32
	Command    Command
	PayloadLen uint32
	Checksum   []byte
}

func BuildHeaderFor(cmd Command, payload []byte) Header {
	return Header{
		MagicWord:  MagicWord,
		Command:    cmd,
		PayloadLen: uint32(len(payload)),
		Checksum:   checksum.DoubleSha256Hash(payload)[0:4],
	}
}

func (h Header) ToBytes() ([]byte, error) {
	header := make([]byte, 24)
	binary.LittleEndian.PutUint32(header[:4], MagicWord)
	// Command (padded with null bytes if less than 12 bytes)
	copy(header[4:16], []byte(h.Command))
	binary.LittleEndian.PutUint32(header[16:20], uint32(h.PayloadLen))
	copy(header[20:24], h.Checksum)
	return header, nil
}
