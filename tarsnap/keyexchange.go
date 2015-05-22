package tarsnap

import (
	"crypto/rand"
	"fmt"

	"github.com/TheCreeper/tarsnap/tarsnap/tarcrypto"
)

/**
 * Connection negotiation and key exchange protocol:
 * Client                                Server
 * Protocol version (== 0; 1 byte)    ->
 *                                    <- Protocol version (== 0; 1 byte)
 * namelen (1 -- 255; 1 byte)         ->
 * User-agent name (namelen bytes)    ->
 *                                    <- 2^x mod p (CRYPTO_DH_PUBLEN bytes)
 *                                    <- RSA-PSS(2^x mod p) (256 bytes)
 *                                    <- nonce (random; 32 bytes)
 * 2^y mod p (CRYPTO_DH_PUBLEN bytes) ->
 * C_auth(mkey) (32 bytes)            ->
 *                                    <- S_auth(mkey) (32 bytes)
 *
 * Both sides compute K = 2^(xy) mod p.
 * Shared "master" key is mkey = MGF1(nonce || K, 48).
 * Server encryption key is S_encr = HMAC(mkey, "S_encr").
 * Server authentication key is S_auth = HMAC(mkey, "S_auth").
 * Client keys C_encr and C_auth are generated in the same way.
 *
 * This is cryptographically similar to SSL where the server has an
 * RSA_DH certificate, except that the client random is omitted (it is
 * unnecessary given that the client provides 256 bits of entropy via
 * its choice owf 2^y mod p).
 */

func (s *Server) BeginKeyExchange() (err error) {

	// Is there a client there?
	b, err := s.Transport.ReadByte()
	if err != nil {

		return
	}

	// Respond back with 0x00
	if b != 0x00 {

		return fmt.Errorf("Client %s did not start connection keyexchange with zero byte\n", s.ConnMetaData.RemoteAddr)
	}
	if err = s.Transport.WriteByte(0x00); err != nil {

		return
	}

	return s.UserAgentLenReceived()
}

func (s *Server) UserAgentLenReceived() (err error) {

	b, err := s.Transport.ReadByte()
	if err != nil {

		return
	}
	if b > MaxUserAgentLen {

		return fmt.Errorf("Client user agent length > %s", MaxUserAgentLen)
	}

	return s.UserAgentReceived()
}

func (s *Server) UserAgentReceived() (err error) {

	b, err := s.Transport.ReadBytes()
	if err != nil {

		return
	}
	if len(b) > MaxUserAgentLen {

		return fmt.Errorf("Client user agent > %s", MaxUserAgentLen)
	}
	s.ConnMetaData.UserAgent = string(b)

	// Generate the keypair
	pub, priv, err := tarcrypto.GenerateSessionKeys()
	if err != nil {

		return
	}

	s.KeyExchange = new(KeyExchange)
	s.KeyExchange.PublicKey = pub
	s.KeyExchange.PrivateKey = priv

	signedBytes, err := tarcrypto.SignBytes(pub, s.ServerKeyChain.PrivateKey)
	if err != nil {

		return
	}

	var a [32]byte
	rand.Read(a[:])

	var p []byte
	p = append(p, pub...)
	p = append(p, signedBytes...)
	p = append(p, a[:]...)

	if err = s.Transport.WriteBytes(p); err != nil {

		return
	}

	return s.DhSent()
}

func (s *Server) DhSent() (err error) {

	return
}
