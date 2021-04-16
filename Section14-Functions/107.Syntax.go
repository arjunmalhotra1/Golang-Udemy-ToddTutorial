package main

import (
	"fmt"
)


func main() {
	foo() // When we call the function we pass in the arguments
	bar("James")
	s1 := woo("MoneyPenny")
	fmt.Println(s1)
	x, y := mouse("Ian","Fleming")
	fmt.Println(x)
	fmt.Println(y)

}

//Syntax of function
// func ( r receiver) identifier(parameters) (return(s)) {
// 	//Code ...
// }
// We define our funtion with parameters

func foo(){
	fmt.Println("Hello from foo")
}

// Everything in Go is "PASS BY VALUE".
// w2hat you see is what you get.
func bar(s string){
	fmt.Println("Hello! ",s)
}

func woo(s string) string{
	return fmt.Sprint("Hello from woo, ",s)
	//Sprint prints to the String.
}

func mouse (fn string, ln string) (string, bool){
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"` )
	b := true
	return a, b
}
