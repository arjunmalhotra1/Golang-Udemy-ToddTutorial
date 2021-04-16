package main

import (
	"fmt"
)


func main() {
	x := sum(2,3,4,5,6,7,8,9)
	fmt.Println("The total  is ",x)
}

// slice of int is accepted here.
func sum(x ...int) int{

	fmt.Println(x)
	fmt.Printf("%T\n",x)
	sum := 0
	// var sum int // This is another way to declare a variable.
	// But sum := 0 is more readable.

	// My solution - This worked
	// for _, v := range x {
	// 	sum = sum + v
	// }

	// His solution
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v,"to the total which is now ", sum)
	}

	return sum
}


// func (r receiver) identifier(parameter(s)) (return(s)) { code }