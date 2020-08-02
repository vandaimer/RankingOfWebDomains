package main

import (
	"fmt"

	"interview/customerimporter"
)

func main() {
	file := "file.csv"
        domains, err := customerimporter.GetListOfDomainsSortedByNumOfOccurrences(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(domains))
}
