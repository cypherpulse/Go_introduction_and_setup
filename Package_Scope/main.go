package main

import "fmt"

var score = 99.5

func main() {
	// In Go, variables declared outside of any function are called package-level variables and are accessible from any function within the same package. This means that you can declare a variable at the package level and use it in multiple functions without having to pass it as an argument.
	fmt.Println("Points:", points)

	sayHello("Alice")
	sayHello("Bob")
	for _, v := range points {
		fmt.Println("Point: ", v)
	}

	showScore()
}
