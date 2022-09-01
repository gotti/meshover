package spec

import (
	"bytes"
	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGenerateKeyPair(t *testing.T) {
	_, err := GenerateKeyPair()
	if err != nil {
		t.Errorf("failed to generate keypair %s", err)
	}
}

func TestNewAddress(t *testing.T) {
	data := []struct {
		testName string
		inIP     net.IP
		out      *Address
	}{
		{
			testName: "1",
			inIP:     net.ParseIP("192.168.0.1"),
			out:      &Address{Ipaddress: &Address_AddressIPv4{AddressIPv4: &AddressIPv4{Ipaddress: "192.168.0.1"}}},
		},
		{
			testName: "2",
			inIP:     net.ParseIP("fd80:dead:beef::1"),
			out:      &Address{Ipaddress: &Address_AddressIPv6{AddressIPv6: &AddressIPv6{Ipaddress: "fd80:dead:beef::1"}}},
		},
	}

	for _, d := range data {
		ad := NewAddress(d.inIP)
		if !proto.Equal(ad, d.out) {
			t.Errorf("NewAddress %s: expected \"%v\", got \"%v\"", d.testName, d.out, ad)
		}
		b := ad.ToNetIP()
		if !b.Equal(d.inIP) {
			t.Errorf("ToNetIP %s: expected \"%v\", got \"%v\"", d.testName, d.inIP, b)
		}
	}
}
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

func TestAddressFormat(t *testing.T) {
	data := []struct {
		testName string
		inIP     net.IP
		out      string
	}{
		{
			testName: "1",
			inIP:     net.ParseIP("192.168.0.1"),
			out:      "192.168.0.1",
		},
		{
			testName: "2",
			inIP:     net.ParseIP("fd80:dead:beef::1"),
			out:      "fd80:dead:beef::1",
		},
	}
	for _, d := range data {
		ap := NewAddress(d.inIP)
		if ap.Format() != d.out {
			t.Errorf("%s: expected \"%s\", got \"%s\"", d.testName, d.out, ap.Format())
		}
	}
}

func TestAddressAndPortFormat(t *testing.T) {
	data := []struct {
		testName string
		inIP     net.IP
		inPort   int32
		out      string
	}{
		{
			testName: "1",
			inIP:     net.ParseIP("192.168.0.1"),
			inPort:   5555,
			out:      "192.168.0.1:5555",
		},
		{
			testName: "2",
			inIP:     net.ParseIP("fd80:dead:beef::1"),
			inPort:   5555,
			out:      "[fd80:dead:beef::1]:5555",
		},
	}

	for _, d := range data {
		ap := &AddressAndPort{Ipaddress: NewAddress(d.inIP), Port: d.inPort}
		if ap.Format() != d.out {
			t.Errorf("%s: expected \"%s\", got \"%s\"", d.testName, d.out, ap.Format())
		}
	}
}

func TestAddressCIDRToAddress(t *testing.T) {
	data := []struct {
		testName string
		inIP     *AddressCIDR
		out      string
	}{
		{
			testName: "1",
			inIP:     &AddressCIDR{Addresscidr: &AddressCIDR_AddressCIDRIPv4{AddressCIDRIPv4: &AddressCIDRIPv4{Mask: 16, Ipaddress: &AddressIPv4{Ipaddress: "192.168.0.1"}}}},
			out:      "192.168.0.1",
		},
		{
			testName: "2",
			inIP:     &AddressCIDR{Addresscidr: &AddressCIDR_AddressCIDRIPv6{AddressCIDRIPv6: &AddressCIDRIPv6{Mask: 64, Ipaddress: &AddressIPv6{Ipaddress: "fd80:dead:beef::1"}}}},
			out:      "fd80:dead:beef::1",
		},
	}

	for _, d := range data {
		ap := d.inIP.ToAddress()
		if ap.Format() != d.out {
			t.Errorf("Address %s: expected \"%s\", got \"%s\"", d.testName, d.out, ap.Format())
		}
		if d.inIP.GetMask() != d.inIP.GetMask() {
			t.Errorf("Mask %s: expected %s, got %s", d.testName, d.out, ap.Format())
		}
	}
}
