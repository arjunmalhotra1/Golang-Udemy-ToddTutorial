package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	// x := func() int {
	// 	return 42
	// }()

	// fmt.Println(x)

	// x := func() int {
	// 	return 42
	// } // Over here we are simply assigning a function to x.
	// fmt.Printf("%T", x)

	x := bar()
	fmt.Printf("%T\n", x)

	// We can run the func() int returned by bar()

	i := x()
	fmt.Println(i)
	//fmt.Println(x()) // This works the same.
}

func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 42
	}
}
