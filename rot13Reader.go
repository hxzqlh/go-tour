package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (n int, err error) {
	temp := make([]byte, 8)
	n, err = rr.r.Read(temp)
	for i := 0; i < n; i++ {
		if (temp[i] >= 'A' && temp[i] <= 'M') || (temp[i] >= 'a' && temp[i] <= 'm') {
			b[i] = temp[i] + 13
		} else if (temp[i] >= 'N' && temp[i] <= 'Z') || (temp[i] >= 'n' && temp[i] <= 'z') {
			b[i] = temp[i] - 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
