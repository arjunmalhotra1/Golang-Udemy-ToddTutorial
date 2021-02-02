package main

import (
	"fmt"
)

func main() {
	x := []int {4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[1])
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) // Does not include index 3. Prints 1 and 2.

	// for i,v := range x {
	// 	fmt.Println(i, v)
	// }

	// for i:=0; i<len(x); i++{
	// 	fmt.Println(x[i:i+1])
	// }

	for i:=0; i<len(x); i++{
		fmt.Println(i, x[i])
	}
}
