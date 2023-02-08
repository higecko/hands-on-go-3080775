// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct{name string}

func (p person) speak() { fmt.Printf("%s speaking ...", p.name)}

func (p person) walk() { fmt.Printf("%s walking ...", p.name)}

// implement the Stringer interface for the person type
func (p person) String() string {
	return fmt.Sprintf("My name is %s", p.name)
}

// define a dog type that can walk but not speak
// type dog struct{name string}

// func (d dog) walk() { fmt.Printf("%s walking ...", d.name)}

func main() {
	// invoke with a person
	p := person{name: "Bob"}
	doHumanThings(p)

	// can we invoke with a dog?
	// doHumanThings(dog{name: "Rufus"})

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
