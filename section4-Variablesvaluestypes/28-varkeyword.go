package main

import (
	"fmt"
)

//x := 42 //This will not work here.

// Here we are declaring a variable "y" and assigns a vlaue 43
// declare and assigning => Initializinzation
var y = 43 // But this WILL work here.

// This declares a variable with IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE at the same time
	x := 42
	fmt.Println("Hello, playground")
	fmt.Println(x)

	// var y = 43
	foo()
	fmt.Println(y)

	//fmt.Println(z)
}

func foo() {
	fmt.Println("y is ", y)
}
