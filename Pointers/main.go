package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func main() {
	name := "tifa"
	updateName(name)

	
	fmt.Println(name)
}
