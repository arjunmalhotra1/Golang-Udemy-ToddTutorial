package main

import (
	"fmt"
)


func main() {
	foo()

	//Anonymous functions
	//aka Anonymous self executing functions
	// This Anonymous func takes no arguments 
	// So we pass no arguments in the end when we do "....}()""
	func(){
		fmt.Println("Anonymous func ran")
	} () // We need to include these "()" paranthesis here as well to 
		// make the anonymous function run/execute.
		// This is what makes the Anonymous func execute.
		// This is to pass in the arguments.

	func (x int){
		fmt.Println("The meaning of life:",x)
	}(42)

}

// This is how we have been looking at the functions so far.
// This is not Anonymous function, it has a name "foo"
func foo() {
	fmt.Println("foo ran") 
}