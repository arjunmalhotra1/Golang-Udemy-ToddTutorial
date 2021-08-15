package main

import "fmt"

// When we want to enclose/scope/contain the scope of a varaible.
// We enclose some code around that variable.
// We are enclosing a variable in some code. So that variable is limited to that area of scope.
var x int //scope of x is the entire package.
func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

}

func foo() {
	fmt.Println("hello")
	x++
}
