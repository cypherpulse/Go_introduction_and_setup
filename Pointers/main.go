package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"
	// updateName(&name)
	fmt.Println("Memeory address of name:", &name)

	m := &name
	// fmt.Println("Memeory address of m:", m)
	// fmt.Println("Value of m:", *m)

	fmt.Println("Before updateName:", name)
	updateName(m)
	fmt.Println("After updateName:", name)

	// fmt.Println(name)
}
