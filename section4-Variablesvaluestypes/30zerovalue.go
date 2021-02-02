package main

import (
	"fmt"
)

var y string
var z int
func main() {
	// DECLARE a variable to be if certaint TYPE
	// and then ASSIGN a V ALUE of that type to the variable
	fmt.Println("printing the value of y",y,"ending")
	fmt.Printf("%T\n", y)
	y = "Bont, James Bond"
	fmt.Println("printing the value of y",y,"ending")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T", z)


}
