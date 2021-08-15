package main

import "fmt"

var x int
var g func()

// var x int = 7
// var g func() = func() {
// 	fmt.Println("g from outseide main")
// }

// We can let the compiler figure it out as well
// var x = 7
// var g = func() {
// 	fmt.Println("g from outside main")
// }

func main() {

	// x := func(i int) {
	// 	fmt.Println(i)
	// }
	// x(42)

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//g()  If we only have in the declaration "var g func()"" then this will give us an error.
	g = f
	g()
	fmt.Printf("This is g %T\n", g)

	fmt.Println("done")
}
