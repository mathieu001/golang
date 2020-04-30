package main

import (
	"fmt"
	"os"
	"strconv"

	"golang/gopl/ch2/ex2.2/conv"
)

func process(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	{
		f := conv.Fahrenheit(t)
		c := conv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))

	}

	{
		f := conv.Feet(t)
		m := conv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", f, conv.FToM(f), m, conv.MToF(m))
	}
	{
		p := conv.Pound(t)
		k := conv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", p, conv.PToK(p), k, conv.KToP(k))
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			process(arg)
		}
		return
	}
	fmt.Println("Input number, Ctrl-C to quit")
	for true {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		process(arg)
	}

}
