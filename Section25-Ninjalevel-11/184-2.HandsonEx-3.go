package main

import "fmt"

type hotdog int

// To show what conversion is
func main() {
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int

	// This doesn't work.
	// y = x // We cannot assign a type hotdog to type int
	// fmt.Println(x)
	// fmt.Printf("%T", x)

	// But this works
	y = int(x) // We are converting type hotdog to type int.
	fmt.Println(y)
	fmt.Printf("%T", y)

}
