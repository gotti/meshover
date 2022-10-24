package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"syscall"

	"github.com/gotti/meshover/frr"
	"github.com/gotti/meshover/internal/meshover"
	"github.com/gotti/meshover/kernel"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
)

//go:embed embed/daemons
var frrDaemonsConfig string

//go:embed embed/vtysh.conf
var frrVtyshConfig string

var (
	controlserver     = flag.String("controlserver", "", "localhost:8080")
	statusserver      = flag.String("statusserver", "", "localhost:8080")
	coilSupport       = flag.Bool("coiladvertise", false, "if set, meshover advertise routes added by coil")
	coilNatSources    = flag.String("coilnatsources", "", "10.128.0.0/16,10.129.0.0/16")
	rawfrrBackend     = flag.String("frr", "", "select one of following: none, dockersdk, nerdctl")
	hostName          = flag.String("hostname", "", "hostname")
	rawRouteGathering = flag.String("gathering", "", "1.1.1.0/27,1.1.2.0/29")
)

var (
	routeGathering []*net.IPNet = nil
	frrBackend     frr.BackendType
)

func getMachineAddresses(isIpv6 bool) (ret []*net.IPNet, err error) {
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for _, l := range links {
		if l.Type() != "device" && l.Type() != "vlan" {
			continue
		}
		var t int
		if isIpv6 {
			t = syscall.AF_INET6
		} else {
			t = syscall.AF_INET
		}
		addr, err := netlink.AddrList(l, t)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch ip address, err=%w", err)
		}
		for _, a := range addr {
			if a.IP.IsGlobalUnicast() {
				ret = append(ret, a.IPNet)
			}
		}
	}
	return ret, nil
}

func getDefaultRouteIF() (ifname *string, err error) {
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for _, l := range links {
		if l.Type() != "device" && l.Type() != "vlan" {
			continue
		}
		addr, err := netlink.AddrList(l, syscall.AF_INET)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch ip address, err=%w", err)
		}
		for _, a := range addr {
			if a.IP.IsGlobalUnicast() {
				return &l.Attrs().Name, nil
			}
		}
	}
	return nil, fmt.Errorf("not found")
}

func parseArgs() error {
	if *hostName == "" {
		h, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("failed to get hostname, err=%w", err)
		}
		hostName = &h
	}
	if *rawRouteGathering != "" {
		for _, r := range strings.Split(*rawRouteGathering, ",") {
			_, n, err := net.ParseCIDR(r)
			if err != nil {
				return fmt.Errorf("failed to parse route gathering option, err=%w", err)
			}
			routeGathering = append(routeGathering, n)
		}
	}
	if *controlserver == "" {
		return fmt.Errorf("controlserver is not specified")
	}
	switch *rawfrrBackend {
	case "none":
		frrBackend = frr.BackendNone
	case "dockersdk":
		frrBackend = frr.BackendDockerSDK
	case "nerdctl":
		frrBackend = frr.BackendNerdCtl
	default:
		return fmt.Errorf("please select valid frr backend with -frr")
	}
	return nil
}

func readConfig(confPath string) string {
	b, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalln("failed to read config")
	}
	return string(b)
}

func panicError(log *zap.Logger, message string, err error) {
	log.Panic(message, zap.Error(err))
}

func createKernelNatSetting(logger *zap.Logger) *kernel.NatSetting {
	if *coilSupport && len(*coilNatSources) > 0 {
		logger.Info("using coil support", zap.String("nat sources", *coilNatSources))
	} else if !*coilSupport {
		logger.Info("coil support disabled")
		return nil
	} else {
		logger.Panic("coilsupport is enabled but nat is disabled. This feature is currently unsupported, please specify nat ip range.")
	}
	ns := []*net.IPNet{}
	for _, s := range strings.Split(*coilNatSources, ",") {
		_, n, err := net.ParseCIDR(s)
		if err != nil {
			panicError(logger, "failed to parse coil nat sources", err)
		}
		ns = append(ns, n)
	}
	destif, err := getDefaultRouteIF()
	if err != nil {
		panicError(logger, "failed to get default route if", err)
	}
	var natsetting *kernel.NatSetting
	if coilSupport != nil && *coilSupport && len(ns) > 0 {
		natsetting = &kernel.NatSetting{
			SourceIPsForNat:   ns,
			DestinationDevice: *destif,
		}
	}
	return natsetting
}

func main() {
	ctx := context.Background()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("failed to initialize zap")
	}
	flag.Parse()
	if err := parseArgs(); err != nil {
		panicError(logger, "failed to parse args", err)
	}
	//get ipv6 unicast address from device or vlan
	addrs, err := getMachineAddresses(true)
	if err != nil {
		panicError(logger, "failed to get ipv6 unicast address", err)
	}
	settings := meshover.Settings{
		HostName:        *hostName,
		ControlServer:   *controlserver,
		StatusServer:    *statusserver,
		UnderlayIP:      addrs[0],
		RouteGathering:  routeGathering,
		FrrVtyshConfig:  frrVtyshConfig,
		FrrDaemonConfig: frrDaemonsConfig,
		FrrBGPConfig:    readConfig("./conf/frr.conf"),
		FrrBackend:      frrBackend,
		KernelSetting: &kernel.Settings{
			CoilSupport: *coilSupport,
			Nat:         createKernelNatSetting(logger),
		},
	}
	if err := meshover.Run(ctx, logger, settings); err != nil {
		panicError(logger, "failed to run meshover", err)
	}
}
