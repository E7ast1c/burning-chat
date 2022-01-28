package pkg

import (
	"crypto/rand"
	"encoding/hex"
)

// Bytes generates n random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

// Hex generates a random hex string with length of hexLength
// e.g: 67aab2d956bd7cc621af22cfb169cba8
const hexLength = 12

func Hex() string {
	return hex.EncodeToString(Bytes(hexLength))
}
