package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/TheCreeper/tarsnap/tarsnap"
)

type ClientConfig struct {

	// Address of the machine to listen on.
	Addr string

	// Location of the keyfile.
	KeyFile string
}

func (cfg *ClientConfig) Validate() (err error) {

	// Check if an addr was specified
	if len(cfg.Addr) == 0 {

		cfg.Addr = tarsnap.DefaultAddr
	}

	// Check root key
	if len(cfg.KeyFile) == 0 {

		return errors.New("Root key path is not specified in config")
	}

	return
}

func GenerateConfig() ([]byte, error) {

	cfg := &ClientConfig{
		Addr:    tarsnap.DefaultAddr,
		KeyFile: "",
	}
	return json.MarshalIndent(cfg, "", "	")
}

func GetCFG() (cfg *ClientConfig, err error) {

	b, err := ioutil.ReadFile(ConfigFile)
	if err != nil {

		return
	}

	if err = json.Unmarshal(b, &cfg); err != nil {

		return
	}

	if err = cfg.Validate(); err != nil {

		return
	}

	return
}
