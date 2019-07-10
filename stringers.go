package main

import "fmt"

type IPAddr [4]byte

//TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ipAdder IPAddr)String() string{
		return fmt.Sprintf("%v.%v.%v.%v",ipAdder[0],ipAdder[1],ipAdder[2],ipAdder[3])
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
