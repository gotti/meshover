package ip

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"net"
)

// SetSeed sets math/random seed
func SetSeed(seed int64) {
	rand.Seed(seed)
}

// GenerateRandomIP generates random IP adddress in specified range.
// example assignAddressRange:
// - 192.168.0.0/16
// - fd00:dead:beef::/48
func GenerateRandomIP(assignAddressRange string) (*net.IPNet, error) {
	i, _, err := net.ParseCIDR(assignAddressRange)
	if err != nil {
		return nil, err
	}
	if v4 := i.To4(); v4 != nil {
		return GenerateRandomIPv4(assignAddressRange)
	} else if v6 := i.To16(); v6 != nil {
		return GenerateRandomIPv6(assignAddressRange)
	}
	return nil, fmt.Errorf("unknown")
}

// GenerateRandomIPv4 generates random IPv4 adddress in specified range.
// example assignAddressRange:
// - 192.168.0.0/16
func GenerateRandomIPv4(assignAddressRange string) (*net.IPNet, error) {
	randomAddr := rand.Uint32()
	i, n, err := net.ParseCIDR(assignAddressRange)
	if err != nil {
		return nil, err
	}
	m, b := n.Mask.Size()
	if b != 32 {
		log.Fatalln("unknown size, you may input ipv6 address range")
	}
	addr := (randomAddr & ((1 << (32 - m)) - 1)) + binary.BigEndian.Uint32(i.Mask(n.Mask))
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, addr)
	return &net.IPNet{IP: ip, Mask: net.CIDRMask(32, 32)}, nil
}

// GenerateRandomIPv6 generates random IPv6 address in specified range. range must be larger than /64
// example assignAddressRange:
// - fd00:dead:beef::/48
func GenerateRandomIPv6(assignAddressRange string) (*net.IPNet, error) {
	randomAddr := rand.Uint64()
	i, n, err := net.ParseCIDR(assignAddressRange)
	if err != nil {
		return nil, err
	}
	m, b := n.Mask.Size()
	if b != 128 {
		log.Fatalln("unknown size, you may input ipv4 address range")
	}
	if m >= 64 {
		log.Fatalln("range must be greater than /64")
	}
	addrUpper := (randomAddr & ((1 << (64 - m)) - 1)) + binary.BigEndian.Uint64(i.Mask(n.Mask))
	addrDowner := uint64(1)
	ip := make(net.IP, 16)
	binary.BigEndian.PutUint64(ip[0:8], addrUpper)
	binary.BigEndian.PutUint64(ip[8:16], addrDowner)
	return &net.IPNet{IP: ip, Mask: net.CIDRMask(64, 128)}, nil
}

// GenerateRandomASN generates random ASN
func GenerateRandomASN() uint32 {
	randomASN := uint32(rand.Intn(94967294)) + uint32(4200000000)
	return randomASN
}
