package main

import (
	"fmt"
	"io"
	"strings"
)

type IOLimitReader struct {
	r io.Reader
	n int64
}

func (iolr *IOLimitReader) Read(p []byte) (n int, err error) {
	n, err = iolr.r.Read(p)
	if n > int(iolr.n) {
		n = int(iolr.n)
	}
	err = io.EOF
	return
}

func LimitReader(r io.Reader, n int64) *IOLimitReader {
	var iolr IOLimitReader
	iolr.r = r
	iolr.n = n
	return &iolr
}

func main() {
	r := LimitReader(strings.NewReader("<html><body><h1>hello</h1></body></html>aaaaaa"), 40)
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	buffer = buffer[:n]
	fmt.Println(n, err, string(buffer))
}
