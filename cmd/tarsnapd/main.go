package main

import (
	"flag"
	"log"
	"os"
)

var (
	// Configuration filepath
	ConfigFile string

	// Configuration generate switch
	GenConfig bool

	// Debug/Verbose switch
	Verbose bool
)

func init() {

	flag.StringVar(&ConfigFile, "f", "", "The configuration file in which the user settings are stored.")
	flag.BoolVar(&GenConfig, "genconf", false, "Generate a configuration and print to stdout.")
	flag.BoolVar(&Verbose, "v", false, "debugging/verbose information")
	flag.Parse()

	if len(ConfigFile) == 0 {

		// Search for config directory
		if len(os.Getenv("XDG_CONFIG_HOME")) != 0 {

			ConfigFile = os.ExpandEnv("$XDG_CONFIG_HOME/tarsnap.conf")

		} else {

			ConfigFile = os.ExpandEnv("$HOME/.config/tarsnap.conf")
		}
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

		return
	}
}

func main() {

	cfg, err := GetCFG()
	if err != nil {

		log.Fatal(err)
	}

	cfg.NewTarsnapListener()
}
