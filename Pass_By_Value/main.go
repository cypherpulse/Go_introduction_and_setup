package main

import (
	"fmt"
)

func updateName(x string){
	x="wedge"
}

func main(){
	// primitive tyes -> they are types that hold a single value, such as int, float, bool, and string. When you pass a primitive type to a function, you are passing a copy of the value. This means that any changes made to the parameter inside the function will not affect the original variable outside the function.

	name:= "tifa"

	updateName(name)
	fmt.Println(name)
}