package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// 改变IPAddr的默认打印格式
func (i IPAddr) String() string {
	lastIndex := len(i) - 1
	var s string
	for i, v := range i {
		s += strconv.Itoa(int(v)) // 数字文字间互转用strconv类
		if i != lastIndex {
			s += "."
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
