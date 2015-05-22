package tarcrypto

import (
	"crypto/rand"

	"github.com/dchest/dhgroup14"
)

func GenerateSessionKeys() (pub, priv []byte, err error) {

	pub, priv, err = dhgroup14.GenerateKeyPair(rand.Reader)
	if err != nil {

		return
	}

	return
}
