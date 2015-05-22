package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/TheCreeper/tarsnap/tarsnap/keyfile"
	"github.com/TheCreeper/tarsnap/tarsnap/tarcrypto"
)

var (

	// Waitgroup for all routines
	wg sync.WaitGroup

	// Debug/Verbose switch
	Verbose bool

	// Root key generate switch
	GenKey string

	// Configuration generate switch
	GenConfig bool

	// Configuration filepath
	ConfigFile string
)

func init() {

	flag.BoolVar(&Verbose, "v", false, "debugging/verbose information")
	flag.StringVar(&GenKey, "k", "", "Generate a root key and write to specified file")
	flag.BoolVar(&GenConfig, "g", false, "Generate a configuration and print to stdout")
	flag.StringVar(&ConfigFile, "f", "", "The configuration file in which the user settings are stored")
	flag.Parse()

	if len(GenKey) != 0 {

		_, priv, err := tarcrypto.GenerateKeyPair()
		if err != nil {

			log.Fatal(err)
		}
		data := tarcrypto.MarshalKey(priv)

		if err := keyfile.Write(GenKey, data); err != nil {

			log.Fatal(err)
		}

		os.Exit(0)
	}

	if GenConfig {

		cfg, err := GenerateConfig()
		if err != nil {

			log.Fatal(err)
		}

		_, err = os.Stdout.Write(cfg)
		if err != nil {

			log.Fatal(err)
		}

		os.Exit(0)
	}

	if len(ConfigFile) == 0 {

		// Search for config directory
		if len(os.Getenv("XDG_CONFIG_HOME")) != 0 {

			ConfigFile = os.ExpandEnv("$XDG_CONFIG_HOME/tarsnap.conf")

		} else {

			ConfigFile = os.ExpandEnv("$HOME/.config/tarsnap.conf")
		}
	}
}

func main() {

	cfg, err := GetCFG()
	if err != nil {

		log.Fatal(err)
	}

	wg.Add(1)
	go cfg.NewTarsnapListener()

	wg.Wait()
}
