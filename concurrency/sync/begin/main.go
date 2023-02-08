// concurrency/sync/begin/main.go
package main

import (
	"sync"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// given a list of names
	names := []string{"James", "Bob", "Karen"}

	// initialize a map to store the length of each name
	namesMap := make(map[string]int)

	// initialize a wait group for the goroutines that will be launched
	var wg sync.WaitGroup
	wg.Add(len(names))

	// launch a goroutine for each name we want to process
	var mu sync.Mutex
	for _, name := range names {
		go func(name string) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			namesMap[name] = len(name)
		}(name)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print the map
	spew.Dump(namesMap)
}
