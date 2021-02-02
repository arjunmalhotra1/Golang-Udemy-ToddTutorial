package main

import (
	"fmt"
)
// creating our won type
type hotdog int
var b hotdog

var a int

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = b // This is not possible
	// Because Go Lang is a static typed language and both
	// a and b are of different type
	// a is of type int nad b is of type hotdog
	// We cant take something of type hotdog and assign it to an int
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
