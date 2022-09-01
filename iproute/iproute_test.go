//go:build maycrashyoursystem
// +build maycrashyoursystem

package iproute

import (
	"os/exec"
	"testing"
)

func TestPermission(t *testing.T){
	o, err := exec.Command("whoami").Output()
	if err != nil {
		t.Errorf("failed to exec whoami")
	}
	t.Error(string(o))
}
