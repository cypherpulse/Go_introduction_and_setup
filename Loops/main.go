package main

import (
	"fmt"
)

func main() {
	//loops are used to execute a block of code repeatedly until a specified condition is met. In Go, there is only one type of loop, the for loop, which can be used in various ways to achieve different looping behaviors.
	// x:= 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of x is ", i)
	// }

	// looping though a slice
	names := []string{"Alice", "Bob", "Charlie", "David"}
	// for index, name := range names {
	// 	fmt.Printf("Index: %d, Name: %s\n", index, name) // %d is for integer and %s is for string
	// }

	for _, name := range names { // if we don't need the index we can use _ to ignore it
		fmt.Printf("Name: %s\n", name)
	}
}
