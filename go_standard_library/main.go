//The standard library in Go provides a wide range of built-in functions and packages that allow developers to perform various tasks without needing to rely on third-party libraries. The standard library includes packages for handling input/output, string manipulation, networking, concurrency, and much more. It is designed to be efficient, reliable, and easy to use, making it an essential part of the Go programming language.

package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){
	greeting := "Hello Bradley loves Go Libraries"

	//strings package - 

	fmt.Println(strings.Contains(greeting,"Bradley"))
	fmt.Println(strings.ReplaceAll(greeting,"Hello","Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"B"))
	fmt.Println(strings.Split(greeting, ""))

	//the original value is unchanged
	fmt.Println("Original string value =",greeting)

	//Sort Package-

	age:=[]int{45,20,35,30,75,60,50,25}
	sort.Ints(ages)s
	fmt.Println(ages)

	index := sort.SearchInts(ages,30)

	names := []strings{"yoshi","mario","peach","bowser","luigi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.println(sort.SearchStrings(names, "bowser"))


}