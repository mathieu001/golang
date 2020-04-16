package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		for _, lines := range strings.Split(string(data), "\r\n") {
			counts[lines]++
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
	// fmt.Printf("%v\n", counts)
}
