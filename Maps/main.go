package main

import (
	"fmt"
)

func main() {
	// In Go, a map is a built-in data type that allows you to store and retrieve values based on a unique key. A map is similar to a dictionary in other programming languages. You can create a map using the make function or by using a map literal.
	// A map literal is a convenient way to create a map with predefined key-value pairs. The syntax for a map literal is map[keyType]valueType{key1: value1, key2: value2, ...}. For example, you can create a map of student names and their corresponding grades like this:

	menu := map[string]float64{
		"Burger": 5.99,
		"Pizza":  8.99,
		"Salad":  4.99,
		"Soda":   1.99,
	}
	fmt.Println(menu)
	fmt.Println("Price of Burger:", menu["Burger"])

	// looping maps
	for k,v:= range menu {
		fmt.Printf("%s: $%.2f\n", k, v)
	}

	// ints as key type
	phoneBook := map[int]string{
		1234567890: "Alice",
		9876543210: "Bob",
		5555555555: "Charlie",
	}
	fmt.Println(phoneBook)
	fmt.Println("Name for 1234567890:", phoneBook[1234567890])

	//updating item in a map
	phoneBook[1234567890] = "David"
	fmt.Println("Updated phone book:", phoneBook)

}
