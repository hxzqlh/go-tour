package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(in byte) byte {
	out := in
	if in >= 'A' && in <= 'M' || in >= 'a' && in <= 'm' {
		out += 13
	} else if in >= 'N' && in <= 'Z' || in >= 'n' && in <= 'z' {
		out -= 13
	}

	return out
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
