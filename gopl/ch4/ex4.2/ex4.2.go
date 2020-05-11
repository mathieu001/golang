package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	encry := flag.String("encry", "sha256", "the encry method")
	flag.Parse()
	// fmt.Println(*encry)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input something")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	switch *encry {
	case "sha256":
		fmt.Printf("sha256=%x\n", sha256.Sum256([]byte(input)))
	case "sha512":
		fmt.Printf("sha512=%x\n", sha512.Sum512([]byte(input)))
	case "sha384":
		fmt.Printf("sha384=%x\n", sha512.Sum384([]byte(input)))
	default:
		fmt.Println("invalid encry")
	}
}
