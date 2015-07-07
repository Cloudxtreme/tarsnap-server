package tarsnap

import (
	"bufio"
	"net"

	"github.com/TheCreeper/tarsnap/tarsnap/keyfile"
	"github.com/TheCreeper/tarsnap/tarsnap/tarcrypto"
)

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

func (s *Server) AddRootKey(f string) {

	data, err := keyfile.Read(f)
	if err != nil {

		panic(err)
	}

	pub, priv, err := tarcrypto.ParseKey(data)
	if err != nil {

		panic(err)
	}

	s.ServerKeyChain = new(ServerKeyChain)
	s.ServerKeyChain.PublicKey = pub
	s.ServerKeyChain.PrivateKey = priv
}

func (s *Server) TriggerCallbacks(p *Packet) (err error) {

	return
}
