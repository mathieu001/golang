package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//read from stdin
		countLines(counts, os.Stdin)
	} else {
		//read from files
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				panic(err)
			}
			countLines(counts, f)
			f.Close()
		}

	}

	//print the map
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(counts map[string]int, f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
