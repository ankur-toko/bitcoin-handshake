package checksum

import "crypto/sha256"

func DoubleSha256Hash(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:]
}
