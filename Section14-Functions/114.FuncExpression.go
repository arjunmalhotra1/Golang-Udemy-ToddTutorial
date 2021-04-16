package main

import (
	"fmt"
)

func main() {
	
	f := func(){
		fmt.Println("My first func expression")
	}

	f() // Since func takes in no parameters we don't ned to pass any
		// arguments to f.
	// Fucntion in GO lang are oftern referred as "Fist calss Citizens"
	// That is, functions can be used as any other variable , like any other int, float,
	// string, slice, struct, map

	g := func(x int){
		fmt.Println("The year Big brother started watching:",x)
	}
	g(1984)
}