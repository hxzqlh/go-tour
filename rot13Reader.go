package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {  
    r io.Reader
}

// ABCDEFG
// HIJKLM
// NOPQRST
// UVWXYZ

func rot13(out byte) byte{
    if out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm' {
        out += 13
    } else if out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z' {
            out -= 13
    }        
    return out
}    

func (fz rot13Reader) Read(b []byte) (int, error){
    n, e := fz.r.Read(b)
    for i := 0; i < n; i++ {
        b[i] = rot13(b[i])    
    }            
    return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}