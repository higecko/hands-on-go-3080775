// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i any = 5
	

	// perform a type assertion and handle the error
	if _, ok := i.(string); !ok {
		fmt.Println("Not an string!")
	} else {
		fmt.Println(i.(string))
	}
}
