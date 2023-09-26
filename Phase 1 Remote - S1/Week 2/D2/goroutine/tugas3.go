package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	input := []string{
		"Portable Network Graphics",
		"As Soon As Possible",
		"Liquid-crystal display",
		"Thank George It's Friday!",
	}

	var wg sync.WaitGroup

	for _, phrase := range input {
		wg.Add(1)
		go func(phrase string) {
			defer wg.Done()
			acronym := generateAcronym(phrase)
			fmt.Printf("%s - %s\n", phrase, acronym)
		}(phrase)
	}

	wg.Wait()
}

func generateAcronym(phrase string) string {
	words := strings.Fields(phrase)
	var acronym string

	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}

	return acronym
}
