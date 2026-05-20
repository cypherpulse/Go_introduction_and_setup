package main

import "fmt"

func main() {
	var name string = "Alice"
	var name2 string = "Bob"
	var name3 string
	fmt.Printf("Hello, %s and %s!\n", name, name2)
	name = "Charlie"
	name2 = "Dave"
	name3 = "Eve"
	fmt.Printf("Hello, %s, %s, and %s!\n", name, name2, name3)

	name4 := "Frank"
	fmt.Printf("Hello, %s!\n", name4)

	//integer
	//ints
	var age int = 30
	var age2 = 25
	age3 := 40
	fmt.Printf("Age: %d, %d, %d\n", age, age2, age3)

	//bits and memory
	// var bigN int8 = 100
	// fmt.Printf("Big number: %d\n", bigN)
	// var bigN2 int16 = 20000
	// fmt.Printf("Big number 2: %d\n", bigN2)
	// var unsignedN uint8 = 255
	// fmt.Printf("Unsigned number: %d\n", unsignedN)

	//float
	var pi float64 = 3.14159
	fmt.Printf("Pi: %.5f\n", pi)
	circleArea := pi * 5 * 5
	fmt.Printf("Area of circle with radius 5: %.2f\n", circleArea)
	var e float32 = 2.71828
	fmt.Printf("Euler's number: %.5f\n", e)
	
}
