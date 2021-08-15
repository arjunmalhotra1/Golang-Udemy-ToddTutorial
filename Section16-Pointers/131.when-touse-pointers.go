package main

import "fmt"

// Pointers are good if you have a large chunk of data and
// you don't want to pass that chunk of data around.
// You just pass the address to where the data is stored.
// Which can save us some performance.
// If we are getting a huge piece of data from the database.
// We can just store it in the memory and pass the address around and use the address wherever we need it.
// Other place we can use the addres is if we want to change something that is at a certain location.

// VIMP
// Everything in Golang is pass by value.
// Drop any concept of Pass by reference or pass by copy. Forget these all rest phrases.
func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
	fmt.Println(x)
}

// func foo(y int) {
// 	fmt.Println(y)
// 	y = 43
// 	fmt.Println(y)
// }

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	//fmt.Println(*y)
	*y = 43
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	//fmt.Println(*y)
}
