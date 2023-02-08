// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Was panicing", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	filename := os.Args[1]
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		panic("FILENAME DOES NOT EXIST")
	}

	// convert the bytes to a string
	contentString := string(contentBytes)

	// initialize a map to store the counts
	const letters string = "letters"
	const symbols string = "symbols"
	const numbers string = "numbers"
	const words string = "words"
	counts := map[string]int{letters: 0, symbols: 0, numbers: 0, words: 0}

	// use the standard library utility package that can help us split the string into words
	contentWords := strings.Fields(contentString)
	
	// capture the length of the words slice
	counts[words] = len(contentWords)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range contentWords {
		for _, char := range word {
			if unicode.IsLetter(char) {
				counts[letters]++
			} else if unicode.IsNumber(char) {
				counts[numbers]++
			} else {
				counts[symbols]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
