package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("error, please input two strings")
		os.Exit(1)
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	if areAnagrams(s1, s2) {
		fmt.Println("are Anagrams")
	} else {
		fmt.Println("not Anagrams")
	}
}

func areAnagrams(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}
	s1 = SortString(s1)
	s2 = SortString(s2)
	if s1 == s2 {
		return true
	} else {
		return false
	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	// s = strings.Join(s, "a") //compiler error
	// return s
	return strings.Join(s, "") //ok
	// ss := strings.Join(s, "") //ok
	// return ss
}
