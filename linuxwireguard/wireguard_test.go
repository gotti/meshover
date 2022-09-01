//go:build maycrashyoursystem
// +build maycrashyoursystem

package linuxwireguard

import (
	"testing"
)

func TestWireguardTunnel(t *testing.T) {
	NewTunnel()
}
