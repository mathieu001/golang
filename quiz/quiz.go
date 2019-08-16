package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"time"
)

func NewShuffle(array [][]string) [][]string {
	for i := len(array) - 1; i >= 0; i-- {
		p := RandInt(i + 1)
		array[i], array[p] = array[p], array[i]
		// fmt.Println("new array=", array)
	}
	return array
}

func RandInt(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)

	wordPtr := flag.String("f", "quiz.csv", "specify the file name")
	boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()
	fmt.Println(*wordPtr)

	csvFile, err := os.Open(*wordPtr)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	totalRows := len(rows)

	fmt.Println("type of rows is", reflect.TypeOf(rows))
	//shuffle the rows if necessary
	if *boolPtr {
		NewShuffle(rows)
		fmt.Println(rows)
	}

	var answer string
	var correct int

	go func() {
		select {
		case s := <-sigs:
			fmt.Println(s)
			fmt.Println("You have got ", correct, "of ", totalRows, "tests")
			os.Exit(0)
		}
	}()

	fmt.Println("Are you ready to finish the quiz with 30 sec.? Press Enter to start")
	fmt.Scanln()

	timer := time.NewTimer(30 * time.Second)
	fmt.Println("timer start")
	go func() {
		<-timer.C
		fmt.Println("timer expired")
		fmt.Println("You have got ", correct, "of ", totalRows, "tests")
		os.Exit(0)
	}()

	for _, row := range rows {
		fmt.Println(row)
		fmt.Println(row[0])
		fmt.Print("Your answer is:\n")
		fmt.Scanln(&answer)
		if answer == row[1] {
			correct++
		}
		fmt.Println()

	}

	fmt.Println("You have got ", correct, "of ", totalRows, "tests")

}
