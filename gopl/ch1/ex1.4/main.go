package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		//read from cmd line
		countLines(os.Stdin, counts, countsFiles)
	} else {
		//read from files
		for _, arg := range files {
			//walk through the files
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 %v\n", err)
				continue
			}
			countLines(f, counts, countsFiles)
			f.Close()
		}
	}

	//printf the map
	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, lines, countsFiles[lines])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsFiles map[string][]string) {
	name := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++

		if !sliceContains(countsFiles[text], name) {
			countsFiles[text] = append(countsFiles[text], name)
		}

	}

}

func sliceContains(a []string, v string) bool {
	for _, item := range a {
		if item == v {
			return true
		}
	}
	return false
}
