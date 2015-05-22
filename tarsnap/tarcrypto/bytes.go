package tarcrypto

import "bytes"

func Compare(a, b []byte) (ok bool) {

	return bytes.EqualFold(a, b)
}
