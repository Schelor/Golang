//stringers.go
package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	s := ""
	for i := 0; i < len(ipAddr); i++ {
		if i == 0 {
			s += strconv.Itoa(int(ipAddr[i]))
		} else {
			s += "." + strconv.Itoa(int(ipAddr[i]))
		}
	}
	return s
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
