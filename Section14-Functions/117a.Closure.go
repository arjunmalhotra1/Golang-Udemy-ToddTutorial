package main

import "fmt"

// When we want to enclose/scope/contain the scope of a varaible.
// We enclose some code around that variable.
// We are enclosing a variable in some code. So that variable is limited to that area of scope.

func main() {
	var x int // If x is declared here then foo() cannot use it.
	fmt.Println(x)
	x++

	{ // these curly braces are enclosing the scope of y
		// Scope of y is only within these curly braces.
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y) // This will give us an error.

	fmt.Println(x)
	//foo()

}

// func foo() {
// 	fmt.Println("hello")
// 	x++
// }
