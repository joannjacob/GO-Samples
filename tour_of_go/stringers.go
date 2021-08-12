package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	convIP := make([]string, len(ip))
	for i, value := range ip {
		convIP[i] = strconv.Itoa(int(value))
	}
	return strings.Join(convIP, ".")
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
