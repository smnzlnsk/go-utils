package utils

import (
	"fmt"
	"net"
)

type Nic struct {
	Name string
	Mtu  int
	Mac  net.HardwareAddr
}

func printInterface(i Nic) {
	fmt.Println(i.Name, i.Mtu, i.Mac)
}
