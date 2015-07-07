package tarsnap

import (
	"crypto/rsa"
	"net"
)

type ServerKeyChain struct {

	// Public key part of the server key par.
	PublicKey *rsa.PublicKey

	// Private key part of the server key pair.
	PrivateKey *rsa.PrivateKey
}

type KeyExchange struct {

	// Public diffie hellman key
	PublicKey []byte

	// Private diffie hellman key
	PrivateKey []byte
}

type ConnMetadata struct {

	// RemoteAddr returns the remote address for this connection.
	RemoteAddr net.Addr

	// LocalAddr returns the local address for this connection.
	LocalAddr net.Addr

	// ClientVersion returns the version of tarsnap used by the client
	UserAgent string
}

type Server struct {

	// Transport
	Transport *Transport

	// Holds server key stuff
	ServerKeyChain *ServerKeyChain

	// Holds data relevent during the exchange process
	KeyExchange *KeyExchange

	// Holds information about the current connected client.
	ConnMetaData ConnMetadata

	// AuthLogCallback, called when a client first connects to the server and trys to authenticate.
	AuthLogCallback func(conn ConnMetadata) error

	// RegisterRequestCallback, called when a machine wants to register with the server.
	RegisterRequestCallback func(conn ConnMetadata, user string) error
}
