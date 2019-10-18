package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-65+13)%26 + 65
	case 'a' <= b && b <= 'z':
		return (b-97+13)%26 + 97
	default:
		return b
	}
}

func (R rot13Reader) Read(b []byte) (int, error) {
	n, err := R.r.Read(b)
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
