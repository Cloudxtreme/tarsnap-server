package tarsnap

import (
	"bufio"
	"crypto/rsa"
	"net"

	"github.com/TheCreeper/tarsnap/tarsnap/keyfile"
	"github.com/TheCreeper/tarsnap/tarsnap/tarcrypto"
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

func NewServer() *Server {

	return new(Server)
}

func (s *Server) ServeConn(conn net.Conn) (err error) {

	s.Transport = &Transport{

		conn: conn,
		r:    bufio.NewReader(conn),
		w:    bufio.NewWriter(conn),
	}
	defer s.Transport.Close()

	s.ConnMetaData = ConnMetadata{

		RemoteAddr: s.Transport.conn.RemoteAddr(),
		LocalAddr:  s.Transport.conn.LocalAddr(),
	}

	if err = s.BeginKeyExchange(); err != nil {

		return
	}

	return
}

func (s *Server) AddRootKey(f string) (err error) {

	data, err := keyfile.Read(f)
	if err != nil {

		return
	}

	pub, priv, err := tarcrypto.ParseKey(data)
	if err != nil {

		return
	}

	s.ServerKeyChain = new(ServerKeyChain)
	s.ServerKeyChain.PublicKey = pub
	s.ServerKeyChain.PrivateKey = priv

	return
}

func (s *Server) TriggerCallbacks(p *Packet) (err error) {

	return
}
