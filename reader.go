type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
    // 赋值并返回
    b[0] = 'A'
    return 1, nil
}

func main() {
    reader.Validate(MyReader{})
}