package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type ClientConfig struct {
	ListenProto string
	ListenAddr  string
	KeyPath     string
}

func (cfg *ClientConfig) Validate() (err error) {

	// Check root key
	if len(cfg.KeyPath) == 0 {

		return errors.New("Root key path is not specified in config")
	}

	return
}

func GenerateConfig() ([]byte, error) {

	cfg := &ClientConfig{}
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
