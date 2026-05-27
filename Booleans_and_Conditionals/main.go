package main

import (
	"fmt"
)

func main() {
	// Booleans and Conditionals
	// age := 45

	// if age < 18 {
	// 	fmt.Println("You are a minor.")
	// } else if age >= 18 && age < 65 {
	// 	fmt.Println("You are an adult.")
	// } else {
	// 	fmt.Println("You are a senior.")
	// }

	// Using boolean and condition in picking scores and grading
	fmt.Println("Enter your score:")
	var score int
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}

}
