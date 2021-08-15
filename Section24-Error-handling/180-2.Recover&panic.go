package main

import "fmt"

// For a real-world example of panic and recover, see the json package from the Go standard library.
// https://golang.org/pkg/encoding/json/
func main() {
	f()
	fmt.Println("Returned notmally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally forom g.") // This didn't get printed.
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
