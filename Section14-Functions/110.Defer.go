package main

import (
	"fmt"
)
// Todd's explanation
// Defer will defer the execution of the fuction until whereever
// it's being called comes to an end.
// Official doc
// A defer statement invokes a function whose execution is deferred
// to the moment the surrounding function returns either becasue the surrounding
// function executed a return statement,
// reached the end of it's fucntion body, or
// because the corresponding goroutine is panicking.

func main() {
	defer foo() // After using "defer" bar() ran first
	// foo() is defered till line 20. As soon as func main() exits then foo() runs.
	// When func() main exits then only the deferred fuction runs.
	// We can use a defer funtion to close a file.
	bar()

}

func foo(){
	fmt.Println("foo")
}
func bar(){
	fmt.Println("bar")
}