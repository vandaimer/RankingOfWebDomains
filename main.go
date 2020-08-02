package main

import (
	"fmt"
	"interview/customerimporter"
	"os"
)

func main() {

	args := os.Args

	if len(args) == 1 {
		fmt.Println("You need to pass the path of the csv file. ./main my-file.csv")
		os.Exit(0)
	}

	file := args[1]

	domains, err := customerimporter.GetListOfDomainsSortedByNumOfOccurrences(file)

	if err != nil {
		os.Exit(0)
	}

	for _, domain := range domains {
		fmt.Printf("Domain: %s | Occurrences: %d\n", domain.Name, domain.Occurrence)
	}
}
