package main

import (
	"fmt"
	"math"
)

// functions are reusable blocks of code that perform a specific task. They allow you to break down your code into smaller, more manageable pieces, and can be called multiple times throughout your program. In Go, you can define a function using the func keyword, followed by the function name, a list of parameters (if any), and the function body enclosed in curly braces.

func sayGreeting(n string) {
	fmt.Printf("Good morning, %s!\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye, %s!\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func areaOfCircle(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Alice")
	// sayBye("Bob")

	// names := []string{"Alice", "Bob", "Charlie", "Bato"}
	// cycleNames(names, sayGreeting)
	// cycleNames(names, sayBye)

	radius := 5.0
	area := areaOfCircle(radius)
	fmt.Printf("The area of a circle with radius %.2f is %.2f\n", radius, area)

}
