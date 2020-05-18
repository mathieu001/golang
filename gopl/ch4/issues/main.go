package main

import (
	"fmt"
	"golang/gopl/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()

	fmt.Println("less than a month")
	for _, item := range result.Items {
		afterOneMonth := item.CreatedAt.AddDate(0, 1, 0)
		if afterOneMonth.After(now) {
			fmt.Printf("%v #%-5d %.55s\n", item.CreatedAt, item.Number, item.Title)
		}
	}
	fmt.Println()

	fmt.Println("less than a year")
	for _, item := range result.Items {
		afterOneYear := item.CreatedAt.AddDate(1, 0, 0)
		if afterOneYear.After(now) {
			fmt.Printf("%v #%-5d %.55s\n", item.CreatedAt, item.Number, item.Title)
		}
	}
	fmt.Println()

	fmt.Println("more than a year")
	for _, item := range result.Items {
		beforeOneYear := item.CreatedAt.AddDate(1, 0, 0)
		if beforeOneYear.Before(now) {
			fmt.Printf("%v #%-5d %.55s\n", item.CreatedAt, item.Number, item.Title)
		}
	}
}
