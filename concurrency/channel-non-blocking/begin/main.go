// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// declare a signal channel to read from
// 	readChan := make(chan struct{})

// 	// use a default case in select to prevent blocking
// 	select {
// 	case <-readChan:
// 		fmt.Println("received from readChan")
// 	default:
// 		fmt.Println("default")
// 	}
// }

func main() {
	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)

	for {
		select {
		case <-timeChan1:
			fmt.Println("timechan1")
		case <-timeChan2:
			fmt.Println("timechan2")
		default:
			fmt.Println("Nothing")
		}
	}
}