package keyfile

import (
	"encoding/pem"
	"errors"
	"io/ioutil"
	"path/filepath"
)

func Read(f string) (data []byte, err error) {

	b, err := ioutil.ReadFile(filepath.Clean(f))
	if err != nil {

		return
	}

	block, _ := pem.Decode(b)
	if block == nil {

		return nil, errors.New("Could not parse PEM data")
	}
	data = block.Bytes

	return
}

func Write(f string, data []byte) (err error) {

	block := &pem.Block{

		Type:  "TARSNAP KEY FILE",
		Bytes: data,
	}
	b := pem.EncodeToMemory(block)

	if err = ioutil.WriteFile(filepath.Clean(f), b, 0600); err != nil {

		return
	}

	return
}
