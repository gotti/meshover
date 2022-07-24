package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gotti/meshover/grpcclient"
	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/wireguard"
	"github.com/vishvananda/netlink"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

var (
	controlserver = flag.String("controlserver", "", "localhost:8080")
	hostName      = flag.String("hostname", "", "hostname")
)

func getMachineAddresses() (ret []*net.IPNet, err error) {
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for _, l := range links {
		if l.Type() != "device" {
			continue
		}
		addr, err := netlink.AddrList(l, syscall.AF_INET6)
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

func generateKeyPair() (privatekey, publickey []byte, err error) {
	if chacha20poly1305.KeySize != curve25519.ScalarSize {
		return nil, nil, fmt.Errorf("assertion failed, key length of chacha20poly1305 and curve25519 mismatch")
	}
	privkey := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(privkey); err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key, err=%w", err)
	}
	pubkey, err := curve25519.X25519(privkey, curve25519.Basepoint)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate public key, err=%w", err)
	}
	return privkey, pubkey, nil
}

/*
func initializeKeyPair() (*spec.KeyPair, error) {
	const filepath = "meshover.conf"
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		privkey, pubkey, err := generateKeyPair()
		if err != nil {
			return nil, fmt.Errorf("failed to generate key pair for authentication, err=%w", err)
		}
		k := new(spec.KeyPair)
		k.PrivateKey = privkey
		k.PublicKey = pubkey
		b, err := proto.Marshal(k)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal keypair, err=%w", err)
		}
		os.WriteFile(filepath, b, 0700)
	}
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read keypair, err=%w", err)
	}
	k := new(spec.KeyPair)
	if err := proto.Unmarshal(b, k); err != nil {
		return nil, fmt.Errorf("failed to unmarshal keypair, err=%w", err)
	}
	return k, nil
}*/

func main() {
	flag.Parse()
	if *hostName == "" {
		h, err := os.Hostname()
		if err != nil {
			log.Fatalf("failed to get hostname, err=%s", err)
		}
		hostName = &h
	}
	if *controlserver == "" {
		log.Fatalln("controlserver is not specified")
	}

	privkey, pubkey, err := generateKeyPair()
	if err != nil {
		log.Fatalf("failed to generate keypair, err=%s\n", err)
	}
	fmt.Printf("pubkey: %x\n", pubkey)

	tunnel := wireguard.NewTunnel("exp-wireguard", privkey, pubkey, *hostName)
	defer tunnel.Close()

	addrs, err := getMachineAddresses()
	if err != nil {
		log.Fatalln(err)
	}
	client, err := grpcclient.NewClient(*hostName, *controlserver, pubkey, addrs[0].IP.String())
	defer client.Close()
	if err != nil {
		log.Fatalln(err)
	}
  tunnel.SetAddress(client.OverlayIP)

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt)

	ticker := time.NewTicker(10 * time.Second)
	watchctx, watchcancel := context.WithCancel(context.Background())
	defer watchcancel()
	go func() {
		for {
			select {
			case <-ticker.C:
				req := new(spec.ListPeersRequest)
				res, err := client.GrpcConn.ListPeers(watchctx, req)
				if err != nil {
					log.Printf("failed to ListPeers err=%s", err)
				}
				ps := res.GetPeers()
				tunnel.SetPeers(ps)
			case <-watchctx.Done():
				return
			}
		}
	}()
	select {
	case <-term:
	case <-tunnel.Dev.Wait():
	}
}
