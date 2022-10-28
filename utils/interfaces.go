package utils

import (
	"fmt"
	"net"
	//v6 "golang.org/x/net/ipv6"
)

type Nic struct {
	Name   string
	Mtu    int
	IPv4   net.Addr
	IPv6   net.Addr
	Status net.Flags
	Mac    net.HardwareAddr
}

func (n Nic) Network() string {
	return n.IPv4.String() + " " + n.IPv6.String()
}

func (n Nic) String() string {
	return n.Name
}

func printInterface(i Nic) {
	fmt.Println(i.Name, i.Mtu, i.Mac)
}
