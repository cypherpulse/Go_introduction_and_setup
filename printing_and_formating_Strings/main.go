package main

import "fmt"

// This program demonstrates how to print and format strings in Go.
func main() {
	age := 25
	name := "Alice"
	fmt.Print("Hello, World!")
	fmt.Print("Welcome to Go programming.")
	fmt.Print("new line using escape key \n")

	//Pritln
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go programming.")
	fmt.Println("my name is:", name, "and my age is:", age)

	//fomatted string
	fmt.Printf("Hello, %s!\n", name)        //here %s is a placeholder for string and \n is for new line
	fmt.Printf("I am %d years old.\n", age) // %d is a placeholder for integer
	//with %v
	fmt.Printf("My name is %v and I am %v years old.\n", name, age) // %v is a placeholder for any value
	//with %q
	fmt.Printf("My name is %q and I am %d years old.\n", name, age) // %q is a placeholder for quoted string
	//with %T
	fmt.Printf("The type of name is %T and the type of age is %T.\n", name, age) // %T is a placeholder for type
	//with %f
	var pi float64 = 3.14159
	fmt.Printf("Pi: %.2f\n", pi) // %.2f is a placeholder for float with 2 decimal places

	// sprintf is for storing formatted string in a variable
	formattedString := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedString)

}
