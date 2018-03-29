package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt rot13Reader) Read(b []byte) (int, error) {
	n, err := rt.r.Read(b)
	for i, v := range b {
		switch {
		case (v >= 65 && v < 78) || (v >= 97 && v < 110):
			b[i] = v + 13
		case (v >= 98 && v < 91) || (v >= 110 && v < 121):
			b[i] = v - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
