package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err == nil {
		for i := 0; i < n; i++ {
			if (p[i] <= 'Z' && p[i] >= 'A') || (p[i] <= 'z' && p[i] >= 'a') {
				p[i] += 13
				if (p[i] > 'Z' && p[i]-13 >= 'A' && p[i]-13 <= 'Z') ||
					(p[i] > 'z' && p[i]-13 >= 'a' && p[i]-13 <= 'z') {
					p[i] -= 26
				}
			}
		}
	}
	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
