package main

import (
	"fmt"
	"log"
	"net"

	"github.com/TheCreeper/tarsnap/tarsnap"
)

func (cfg *ClientConfig) NewTarsnapListener() {

	server := tarsnap.NewServer()
	server.RegisterRequestCallback = RegisterRequestHandler

	// Import root key
	if err := server.AddRootKey(cfg.KeyPath); err != nil {

		log.Print(err)
	}

	// Listen on TCP port 2000 on all interfaces.
	ln, err := net.Listen(tarsnap.DefaultProto, tarsnap.DefaultListenAddr)
	if err != nil {

		log.Fatal(err)
	}
	defer ln.Close()

	for {

		conn, err := ln.Accept()
		if err != nil {

			log.Print(err)
		}
		defer conn.Close()

		if err := server.ServeConn(conn); err != nil {

			log.Print(err)
		}
	}
}

func AuthLogCallbackHandler(conn tarsnap.ConnMetadata) (err error) {

	log.Printf("New connection attempt from %s with client version %s\n",
		conn.RemoteAddr,
		conn.UserAgent)

	return
}

func RegisterRequestHandler(conn tarsnap.ConnMetadata, user string) (err error) {

	if len(user) >= tarsnap.MaxUserLen {

		return fmt.Errorf("User name too long: %s", user)
	}

	return
}
