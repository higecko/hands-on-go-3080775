// concurrency/channels/begin/main.go
package main

import (
	"fmt"
	// "time"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<-int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)
	// invoke the sum function as a goroutine
	go sum(nums, ch)
	result := <-ch
	fmt.Println("Result", result)

	ch2 := make(chan string)
	ch2 <- "James"
	ch2 <- "Bob"
	fmt.Println(<-ch2, <-ch2)
}