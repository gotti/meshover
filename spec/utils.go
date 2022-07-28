package spec

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

func GenerateKeyPair() (key *Curve25519KeyPair, err error) {
	if chacha20poly1305.KeySize != curve25519.ScalarSize {
		return nil, fmt.Errorf("assertion failed, key length of chacha20poly1305 and curve25519 mismatch")
	}
	privkey := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(privkey); err != nil {
		return nil, fmt.Errorf("failed to generate private key, err=%w", err)
	}
	pubkey, err := curve25519.X25519(privkey, curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("failed to generate public key, err=%w", err)
	}
  publickey := new(Curve25519Key)
  publickey.Key = pubkey
  privatekey := new(Curve25519Key)
  privatekey.Key = privkey
  ret := new(Curve25519KeyPair)
  ret.Publickey = publickey
  ret.Privatekey = privatekey
	return ret, nil
}

func (c *Curve25519Key)EncodeBase64() string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(c.Key)))
	base64.StdEncoding.Encode(dst, c.Key)
	return string(dst)
}

func (a *AddressAndPort)Format() string{
  i := a.GetIpaddress().GetIpaddress()
  p := a.GetPort()
  n := net.ParseIP(i)
  if n.To4() != nil {
    return fmt.Sprintf("%s:%d", i, p)
  }
  if n.To16() != nil {
    return fmt.Sprintf("[%s]:%d", i, p)
  }
  return ""
}

func NewAddress(n net.IP) *Address{
  return &Address{Ipaddress: n.String()}
}

func (a *Address)ToNetIP() net.IP{
  n := net.ParseIP(a.GetIpaddress())
  return n
}
