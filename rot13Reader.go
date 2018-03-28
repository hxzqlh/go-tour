package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot(b byte) byte {
	b = b + 1
	return b
}
func (ro rot13Reader) Read(b []byte) (int, error) {
	n, err := ro.r.Read(b)
	fmt.Println(n, err) //此时n为读取的字节数，err为nil
	for i := 0; i < n; i++ {
		b[i] = rot(b[i])
	}
	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") //获得一个带缓冲的reader
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
