// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}

func count(s string, f func(r rune) bool) int {
	c := 0
	for _, char := range s {
		if f(char) {
			c++
		}
	}
	return c
}

type letterCounter struct{ identifier string }
func (lc letterCounter) name() string { return lc.identifier}
func (lc letterCounter) count(input string) int {
	return count(input, unicode.IsLetter)
}

type numberCounter struct{ designation string }
func (nc numberCounter) name() string { return nc.designation}
func (nc numberCounter) count(input string) int {
	return count(input, unicode.IsNumber)
}

type symbolCounter struct{ label string }
func (sc symbolCounter) name() string { return sc.label}
func (sc symbolCounter) count(input string) int {
	return count(input, func(r rune) bool {return !(unicode.IsLetter(r) || unicode.IsNumber(r))})
}

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	// loop over the counters and use their name as the key
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}

	// return the map
	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)
	spew.Dump(data)

	// call doAnalysis and pass in the data and the counters
	lc := letterCounter{identifier: "letters"}
	nc := numberCounter{designation: "numbers"}
	sc := symbolCounter{label: "symbols"}
	analysis := doAnalysis(data, lc, nc, sc)

	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
