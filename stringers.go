package main

import (
	"fmt"
	//	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var a string
	for _, v := range ip {
		var s string
		s = fmt.Sprintf("%v", v)
		a = a + s
		a = a + "."
	}
	return strings.Trim(a, ".")
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
