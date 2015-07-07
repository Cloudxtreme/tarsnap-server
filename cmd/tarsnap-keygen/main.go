package main

import (
	"flag"
	"log"

	"github.com/TheCreeper/tarsnap/tarsnap/keyfile"
	"github.com/TheCreeper/tarsnap/tarsnap/tarcrypto"
)

var (
	// Location of the private key.
	KeyFile string

	// Generate private key switch.
	GenKey bool

	// Generate the crypto_keys_server.c file for tarsnap.
	GenSource bool

	// Debug/Verbose switch
	Verbose bool
)

func init() {
	flag.StringVar(&KeyFile, "f", "", "Location of the keyfile to write the generated key to.")
	flag.BoolVar(&GenKey, "k", false, "Generate a root key and write to specified file.")
	flag.BoolVar(&GenSource, "s", false, "Generate the crypto_keys_server.c file.")
	flag.BoolVar(&Verbose, "v", false, "debugging/verbose information")
	flag.Parse()
}

func main() {
	if GenKey {
		_, priv, err := tarcrypto.GenerateKeyPair()
		if err != nil {
			log.Fatal(err)
		}
		data := tarcrypto.MarshalKey(priv)

		if err := keyfile.Write(KeyFile, data); err != nil {
			log.Fatal(err)
		}
		return
	}

	if GenSource {
		b, err := keyfile.Read(KeyFile)
		if err != nil {
			log.Fatal(err)
		}

		if err := keyfile.GenerateSource(b); err != nil {
			log.Fatal(err)
		}
		return
	}
}
