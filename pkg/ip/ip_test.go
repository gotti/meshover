package ip

import (
	"net"
	"testing"
)

func TestGenerateRandomIPv4(t *testing.T) {
	data := []struct {
		name   string
		prefix string
	}{
		{
			name:   "1",
			prefix: "192.168.0.0/16",
		},
	}

	for _, d := range data {
		i, err := GenerateRandomIPv4(d.prefix)
		if err != nil {
			t.Errorf("error happend TestData=%s, Prefix=%s, error=%s\n", d.name, d.prefix, err)
		}
		_, n, err := net.ParseCIDR(d.prefix)
		if err != nil {
			t.Errorf("invalid CIDR, error=%s", err)
		}
		if !n.Contains(i.IP) {
			t.Errorf("generated ip is out of prefix, generated=%s, prefix=%s", i, d.prefix)
		}
	}
}

func TestGenerateRandomIPv6(t *testing.T) {
	data := []struct {
		name   string
		prefix string
	}{
		{
			name:   "1",
			prefix: "fd00::/56",
		},
	}

	for _, d := range data {
		i, err := GenerateRandomIPv6(d.prefix)
		if err != nil {
			t.Errorf("error happend TestData=%s, Prefix=%s, error=%s\n", d.name, d.prefix, err)
		}
		_, n, err := net.ParseCIDR(d.prefix)
		if err != nil {
			t.Errorf("invalid CIDR, error=%s", err)
		}
		if !n.Contains(i.IP) {
			t.Errorf("generated ip is out of prefix, generated=%s, prefix=%s", i, d.prefix)
		}
	}
}
