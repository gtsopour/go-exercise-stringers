package main

import (
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
// Dynamic length solution
func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	for i, value := range ip {
		// strconv.Itoa() convers int to string
		s[i] = strconv.Itoa(int(value))
	}	
	return strings.Join(s, ".")
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
