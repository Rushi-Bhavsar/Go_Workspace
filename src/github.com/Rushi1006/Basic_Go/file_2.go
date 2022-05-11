package main

import "fmt"

// Not supported in global declaration
//b := 1

// Support using var keyword.
var b = 1

func main() {
	// Shorthand declaration operator.
	a := 1
	fmt.Println(a)
	a = 2
	fmt.Println(a)
	var c string
	var d int
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
