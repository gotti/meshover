package kernel

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// Instance is kernel instance
type Instance interface {
}

// NatSetting is setting to enable nat, for miscellaneous kubernetes pods
type NatSetting struct {
	SourceIPsForNat   []*net.IPNet
	DestinationDevice string
}

// Settings have kernel Parameter Settings
type Settings struct {
	//CoilSupport defines wheather creating coil vrf
	CoilSupport bool
	Nat         *NatSetting
}

// InstanceKernel implements Instance
type InstanceKernel struct {
	settings *Settings
}

func execCommand(command string, args ...string) error {
	res, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("command execution failed, output=%s, err=%w", string(res), err)
	}
	return nil
}

// NewKernelInstance creates instance and apply initial settings
func NewKernelInstance(settings *Settings) (*InstanceKernel, error) {
	if settings == nil {
		return nil, nil
	}
	if err := execCommand("sysctl", "-w", "net.ipv4.ip_forward=1"); err != nil {
		return nil, fmt.Errorf("failed to enable ipv4 forwarding")
	}
	if err := execCommand("sysctl", "-w", "net.ipv4.conf.all.rp_filter=0"); err != nil {
		return nil, fmt.Errorf("failed to disable all rp_filter, err=%w", err)
	}
	if err := execCommand("sysctl", "-w", "net.ipv4.conf.default.rp_filter=0"); err != nil {
		return nil, fmt.Errorf("failed to disable default rp_filter, err=%w", err)
	}
	if settings.CoilSupport {
		for _, c := range settings.Nat.SourceIPsForNat {
			if err := execCommand("iptables", "-t", "nat", "-A", "POSTROUTING", "-o", settings.Nat.DestinationDevice, "-s", c.String(), "-j", "MASQUERADE"); err != nil {
				return nil, fmt.Errorf("failed to enable natting")
			}
		}
	}
	if settings.CoilSupport {
		if err := execCommand("ip", "link", "add", "coilex", "type", "vrf", "table", "119"); err != nil {
			if !strings.Contains(err.Error(), "File exists") {
				return nil, fmt.Errorf("failed to set vrf")
			}
		}
	}
	return &InstanceKernel{settings: settings}, nil
}

// InstanceDummy is a dummy instance do nothing
type InstanceDummy struct {
}

// NewDummyInstance creates Dummy Instance
func NewDummyInstance() *InstanceDummy {
	return &InstanceDummy{}
}
