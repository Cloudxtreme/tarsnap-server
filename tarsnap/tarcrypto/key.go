package tarcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

func GenerateKeyPair() (pub *rsa.PublicKey, priv *rsa.PrivateKey, err error) {

	pKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {

		return
	}
	if err = pKey.Validate(); err != nil {

		return
	}

	priv = pKey
	pub = &pKey.PublicKey

	return
}

func ParseKey(b []byte) (pub *rsa.PublicKey, priv *rsa.PrivateKey, err error) {

	pKey, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {

		return
	}
	pKey.Precompute()

	priv = pKey
	pub = &pKey.PublicKey

	return
}

func MarshalKey(priv *rsa.PrivateKey) []byte {

	return x509.MarshalPKCS1PrivateKey(priv)
}
