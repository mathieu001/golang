package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	go mustCopy(os.Stdout, c)
	mustCopy(c, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
