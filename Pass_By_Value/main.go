package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["Yoghurt"] = 8.99
}

func main() {
	// primitive tyes -> they are types

	name := "tifa"

	name = updateName(name)
	fmt.Println(name)

	// Derived types -> they are references to the underlying data structure. When you assign a non-primitive type to a new variable, you are creating a reference to the same underlying data structure, rather than creating a new copy of the data. This means that changes made to the data through one variable will affect all variables that reference the same data.
	// In Go, the following types are considered derived types: arrays, slices, maps, structs, channels, and interfaces. When you assign a variable of one of these types to another variable, you are creating a reference to the same underlying data structure. For example:
	menu := map[string]float64{
		"Burger": 5.99,
		"Pizza":  7.99,
		"Salad":  4.99,
		"Soda":   1.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
