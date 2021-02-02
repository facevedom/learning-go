package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/* Implement `WordCount`. It should return a map of the counts of each "word" 
   in the string `s`. The `wc.Test` function runs a test suite against the
   provided function and prints success or failure.
   
   You might fing strings.Fields (https://golang.org/pkg/strings/#Fields)
   helpful.
   
   Fields splits the string `s` around each instance of one or more consecutive
   white space characters, as defined by unicode.IsSpace, returning a slice of
   substrings of `s` or an empty slice if `s` contains only white space. */
func WordCount(s string) map[string]int {
	ocurrences := make(map[string]int)

	// Split the string into words
	words := splitIntoWords(s)

	for _, word := range words {
		// If the word is in the map, add 1 to the ocurrences
		if _, present := ocurrences[word]; present {
			ocurrences[word] = ocurrences[word] + 1
		} else {
			// else add the word as key and 1 as occurrence
			ocurrences[word] = 1
		}
	}
	return ocurrences
}

func splitIntoWords(s string) []string {
	words := strings.Fields(s)
	return words
}

func main() {
	wc.Test(WordCount)
}
