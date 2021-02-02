package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12) // Type is int, length is 10 and capacity is 100
	// capacity 12/100 means we have 12/100 spots in the underlying array to use.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// This works
	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	// But his will NOT work as index 10 is out of the range of the slice
	// x[10] = 3243
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// But what we can do is this
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

