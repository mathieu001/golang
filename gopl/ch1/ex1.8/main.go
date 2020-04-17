package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//check prefix
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		//use io.copy
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}

		resp.Body.Close()
	}
}
