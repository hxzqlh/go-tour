package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func Rot13(x byte) byte {
	switch {
	case x >= 'A' && x <= 'M':
		return x + 13
	case x >= 'a' && x <= 'm':
		return x + 13
	case x >= 'N' && x <= 'Z':
		return x - 13
	case x >= 'n' && x <= 'z':
		return x - 13
	default:
		return x
	}
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	tmp := make([]byte, 16)
	n, err := rr.r.Read(tmp)
	for i := 0; i < n; i++ {
		b[i] = Rot13(tmp[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
