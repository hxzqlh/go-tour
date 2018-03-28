package main

import (
	"io"
	"os"
	"strings"
)

// A-65 M-77  N-78 Z-90
// a-97 m-109 n-110 z-122

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b <= 77 && b >= 65 || b <= 109 && b >= 97 {
		b += 13
	} else if b <= 90 && b >= 78 || b <= 122 && b >= 110 {
		b -= 13
	}
	return b
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
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
