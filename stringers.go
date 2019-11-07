package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法string(add)
func (addr IPAddr) String() string {
	strAddr := make([]string, 4)
	for i, add := range addr {
		strAddr[i] = strconv.Itoa(int(add))
	}
	return strings.Join(strAddr, ".")
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
