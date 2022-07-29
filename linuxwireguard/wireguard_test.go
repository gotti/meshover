package linuxwireguard

import (
	"net"
	"testing"
)

func TestGenerateLinkLocalV6(t *testing.T) {
	i := net.ParseIP("192.168.1.1")
	if generateLinkLocalV6(i).String() != "fe80:0:dead:beef:dead:beef:c0a8:101" {
		t.Errorf("expected fe80:0:dead:beef:0a1f:ca6:c0a8:0101, got=%s", generateLinkLocalV6(i).String())
	}
}
