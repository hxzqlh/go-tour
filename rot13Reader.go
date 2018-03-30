package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	length, err := r13.r.Read(b)

	if err != nil {
		return length, err
	}

	for i := 0; i < length; i++ {
		v := b[i]

		switch {
		case 'a' <= v && v <= 'm':
			fallthrough
		case 'A' <= v && v <= 'M':
			b[i] += 13
		case 'n' <= v && v <= 'z':
			fallthrough
		case 'N' <= v && v <= 'Z':
			b[i] -= 13
		}
	}
	return length, nil
}
