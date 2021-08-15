package main

import "fmt"

func main() {
	a := incrementor() // a is the function returned by the incrementor.
	b := incrementor() // b is a different variable than a.
	fmt.Println(a())   // Values of a and b are stored in the different memory locations.
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {

	var x int
	return func() int {
		x++ // x is available inside this function too.
		// Without having us to pass x as an argument into this function func(x)
		return x
	}
}
