package message

type VerAck struct {
}

func (v VerAck) ToBytes() ([]byte, error) {
	return []byte{}, nil
}
