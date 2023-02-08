// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

type printable interface {
	~int | ~string | bool
}
func print[T printable](value T) { fmt.Println(value)}

// Part 2 sum function refactoring

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

type numeric interface {
	constraints.Integer | constraints.Float
}
func sumAny[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	// call non-generic print functions
	printString("Hello")
	printInt(42)
	printBool(true)

	// call generic printAny function for each value above
	print("Hello")
	print(42)
	print(true)

	// call sum function
	fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	fmt.Println("resultAny", sumAny(1, 2, 3))
	fmt.Println("resultAny", sumAny(1.2, 2.3, 3.4))
}
