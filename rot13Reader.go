package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
//代换密码
func rot13(b byte) byte{
	switch{
	case b>'A'&&b<'Z':
		b = b+8;	
	case b>'a'&&b<'z':
		b = b-8;
	}
	return b
}

//实现Read方法
func (rotR rot13Reader)Read(b []byte)(int,error){
	n,err := rotR.r.Read(b)
	for i:= 0;i<n;i++{
		b[i] = rot13(b[i])
	}
	return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
