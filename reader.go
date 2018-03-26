package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	var err error
	return 1, err
}

func main() {
	reader.Validate(MyReader{})
}