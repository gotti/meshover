package dummy

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

//Device shows Dummy Interface
type Device struct{
	link netlink.Link
}

//NewDevice creates a Dummy device
func NewDevice(name string) (*Device, error){
	attr := netlink.NewLinkAttrs()
	attr.Name = name
	dum := &netlink.Dummy{
		LinkAttrs: attr,
	}
	if err := netlink.LinkAdd(dum); err != nil {
		return nil, fmt.Errorf("failed to create dummy interface %w", err)
	}
	if err := netlink.LinkSetUp(dum); err != nil {
		return nil, fmt.Errorf("failed to up dummy interface %w", err)
	}
	return &Device{link: dum}, nil
}

//AddAddress adds ip address to dummy device
func (d *Device)AddAddress(n *net.IPNet) error{
	addr, err := netlink.ParseAddr(n.String())
	if err != nil {
		return fmt.Errorf("failed to parse IPNet, err=%w", err)
	}
	if err := netlink.AddrAdd(d.link, addr); err != nil {
		return fmt.Errorf("failed to add address %w", err)
	}
	return nil
}

//Clean deletes interface
func (d *Device)Clean() error{
	return netlink.LinkDel(d.link)
}
