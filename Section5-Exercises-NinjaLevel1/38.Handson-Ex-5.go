package main

import (
	"fmt"
)
type hotdog int
var x hotdog
var y int
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
	y = int(x) // What I had done was "y:= int(x)"
	// doing := beats the purpose of the exercise
	fmt.Println(y)
	fmt.Printf("%T\n",y)
}
