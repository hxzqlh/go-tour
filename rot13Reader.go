package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (ro rot13Reader) Read(buf []byte) (int, error) {
	length, err := ro.r.Read(buf)
	if err != nil {
		return length, err
	}

	for i := 0; i < length; i++ {
		v := buf[i]
		switch {
		case 'a' <= v && v <= 'm':
			fallthrough
		case 'A' <= v && v <= 'M':
			buf[i] = v + 13
		case 'n' <= v && v <= 'z':
			fallthrough
		case 'N' <= v && v <= 'Z':
			buf[i] = v - 13
		}
	}
	return length, nil

}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
