package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	ret := ""
	for _, i := range ip {
		ret += strconv.Itoa(int(i)) + "."
	}
	return ret[:len(ret)-1]
}


func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}