package tarcrypto

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Sum(r io.Reader) (hashed []byte, err error) {

	hasher := sha256.New()

	_, err = io.Copy(hasher, r)
	if err != nil {

		return
	}

	hashed = make([]byte, sha256.BlockSize)
	hex.Encode(hashed, hasher.Sum(nil))

	return
}
