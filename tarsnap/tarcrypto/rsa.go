package tarcrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func SignBytes(data []byte, priv *rsa.PrivateKey) (s []byte, err error) {

	hash := sha256.Sum256(data)
	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hash[:], &rsa.PSSOptions{})
}
