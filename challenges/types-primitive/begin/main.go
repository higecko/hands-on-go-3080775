// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 123

	// print the value of the local variable "val"
	fmt.Println("local val", val)
	// print the value of the package-level variable "val"
	printGlobalVal()
	// update the package-level variable "val" and print it again
	updateGlobalVal()
	printGlobalVal()

	// print the pointer address of the local variable "val"
	fmt.Println("local pointer", &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	valAddress := &val
	*valAddress = 900
	fmt.Println("new val", val)

}

func printGlobalVal() {
	fmt.Println("global val", val)
}

func updateGlobalVal() {
	val = "not global"
}
