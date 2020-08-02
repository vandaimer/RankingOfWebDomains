// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Domain struct {
	Name       string
	Occurrence int
}

func GetListOfDomainsSortedByNumOfOccurrences(fileName string) ([]Domain, error) {
	lines, err := readLines(fileName)

	if err != nil {
		fmt.Println("Error on loading the file")
		return nil, err
	}

	domains := getListOfDomains(lines)

	occurrencesDomains := countOccurrencesForEachDomain(domains)

	return sortByDomain(occurrencesDomains), nil
}

func countOccurrencesForEachDomain(domains []string) map[string]int {
	domainsOccurrences := map[string]int{}

	for _, domain := range domains {
		// Only as a document, I know that when the map does not have an entry, the return will be zero
		// So, it's safe to increment the value without check if its exists
		domainsOccurrences[domain]++
	}

	return domainsOccurrences
}

func getListOfDomains(lines [][]string) []string {
	domains := make([]string, 0, len(lines))

	for _, item := range lines[1:] {
		email := item[2]
		if domain, err := getDomainFromEmail(email); err == nil {
			domains = append(domains, domain)
		}
	}

	return domains
}

func getDomainFromEmail(email string) (string, error) {
        // Here we can improve to use regex
        // I'll keep in this way for the first version of this code

	s := strings.Split(email, "@")

	if len(s) == 1 {
		return "", errors.New("Invalid email could not be parsed.")
	}

	return s[1], nil
}

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rows := [][]string{}

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func sortByDomain(domains map[string]int) []Domain {
	sortedMap := make([]Domain, 0, len(domains))
	onlyDomains := make([]string, 0, len(domains))

	for domain := range domains {
		onlyDomains = append(onlyDomains, domain)
	}

	sort.Strings(onlyDomains)

	for _, domain := range onlyDomains {
		sortedMap = append(sortedMap, Domain{
			domain,
			domains[domain],
		})
	}

	return sortedMap
}
