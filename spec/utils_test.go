package spec

import (
	"bytes"
	"net"
	"testing"
)

func TestNewAddressCIDR(t *testing.T) {
	data := []struct {
		name string
		in   string
	}{
		{
			name: "v4",
			in:   "192.168.0.1/16",
		},
		{
			name: "v6",
			in:   "fd80::/56",
		},
	}
	for _, d := range data {
		_, n, err := net.ParseCIDR(d.in)
		if err != nil {
			t.Errorf("invalid cidr %s", d.in)
		}
		addr := NewAddressCIDR(*n)
		n2 := addr.ToNetIPNet()
		if !(n.IP.Equal(n2.IP) && bytes.Compare(n.Mask, n2.Mask) == 0) {
			t.Errorf("error %s, %s", n, n2)
		}
	}
}
