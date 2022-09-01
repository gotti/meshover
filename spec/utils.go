package spec

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"strconv"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

// GenerateKeyPair generates curve25519 keypair
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

// EncodeBase64 encodes curve25519 key
func (c *Curve25519Key) EncodeBase64() string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(c.Key)))
	base64.StdEncoding.Encode(dst, c.Key)
	return string(dst)
}

// Format formats AddressAndPort
func (a *AddressAndPort) Format() string {
	var n net.IP
	if v4 := a.GetIpaddress().GetAddressIPv4(); v4 != nil {
		n = net.ParseIP(v4.GetIpaddress())
	} else if v6 := a.GetIpaddress().GetAddressIPv6(); v6 != nil {
		n = net.ParseIP(v6.GetIpaddress())
	}
	p := a.GetPort()
	if n.To4() != nil {
		return fmt.Sprintf("%s:%d", n, p)
	}
	if n.To16() != nil {
		return fmt.Sprintf("[%s]:%d", n, p)
	}
	return ""
}

// NewAddressCIDR creates AddressCIDR from [net.IPNet]
func NewAddressCIDR(n net.IPNet) *AddressCIDR {
	m, _ := n.Mask.Size()
	if r := n.IP.To4(); r != nil {
		return &AddressCIDR{Addresscidr: &AddressCIDR_AddressCIDRIPv4{&AddressCIDRIPv4{Ipaddress: &AddressIPv4{Ipaddress: n.IP.String()}, Mask: int32(m)}}}
	} else if r := n.IP.To16(); r != nil {
		return &AddressCIDR{Addresscidr: &AddressCIDR_AddressCIDRIPv6{&AddressCIDRIPv6{Ipaddress: &AddressIPv6{Ipaddress: n.IP.String()}, Mask: int32(m)}}}
	}
	return nil
}

// NewAddress creates Address from [net.IP]
func NewAddress(n net.IP) *Address {
	if r := n.To4(); r != nil {
		return &Address{Ipaddress: &Address_AddressIPv4{&AddressIPv4{Ipaddress: n.String()}}}
	} else if r := n.To16(); r != nil {
		return &Address{Ipaddress: &Address_AddressIPv6{&AddressIPv6{Ipaddress: n.String()}}}
	}
	return nil
}

// ToNetIP translates from Address into [net.IP]
func (a *Address) ToNetIP() net.IP {
	if v4 := a.GetAddressIPv4(); v4 != nil {
		r := net.ParseIP(v4.GetIpaddress())
		return r
	}
	if v6 := a.GetAddressIPv6(); v6 != nil {
		r := net.ParseIP(v6.GetIpaddress())
		return r
	}
	return nil
}

// Format formats Address to like "192.168.0.0"
func (a *Address) Format() string {
	return string(a.ToNetIP().String())
}

// ToNetIPNet translates from AddressCIDR into [net.IPNet]
func (n *AddressCIDR) ToNetIPNet() *net.IPNet {
	if v4 := n.GetAddressCIDRIPv4(); v4 != nil {
		r := net.ParseIP(v4.GetIpaddress().GetIpaddress())
		te := net.IPNet{IP: r, Mask: net.CIDRMask(int(v4.GetMask()), 32)}
		return &te
	}
	if v6 := n.GetAddressCIDRIPv6(); v6 != nil {
		r := net.ParseIP(v6.GetIpaddress().GetIpaddress())
		return &net.IPNet{IP: r, Mask: net.CIDRMask(int(v6.GetMask()), 128)}
	}
	return nil
}

// ToAddress converts AddressCIDR to Address
func (n *AddressCIDR) ToAddress() *Address {
	var ret Address
	if v4 := n.GetAddressCIDRIPv4(); v4 != nil {
		r := v4.GetIpaddress()
		ret.Ipaddress = &Address_AddressIPv4{AddressIPv4: r}
		return &ret
	}
	if v6 := n.GetAddressCIDRIPv6(); v6 != nil {
		r := v6.GetIpaddress()
		ret.Ipaddress = &Address_AddressIPv6{AddressIPv6: r}
		return &ret
	}
	return nil
}

// GetMask returns mask
func (n *AddressCIDR) GetMask() int32 {
	if v4 := n.GetAddressCIDRIPv4(); v4 != nil {
		return v4.GetMask()
	}
	if v6 := n.GetAddressCIDRIPv6(); v6 != nil {
		return v6.GetMask()
	}
	return -1
}

// Format formats to string like "192.168.0.0/16"
func (n *AddressCIDR) Format() string {
	return n.ToAddress().Format() + "/" + strconv.Itoa(int(n.GetMask()))
}

// Format returns ASN with string
func (a *ASN) Format() string {
	return strconv.FormatUint(uint64(a.GetNumber()), 10)
}
