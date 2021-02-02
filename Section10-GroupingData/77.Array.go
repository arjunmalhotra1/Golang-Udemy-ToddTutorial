package main

import (
	"fmt"
)

func main() {
	var x [5] int // Single type elements
	var y [6] int // This is totally different type as length,
				  // of an array is partof type of the array.
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))


}
