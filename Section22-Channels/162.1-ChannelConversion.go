package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	// General to specific converts
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// General to specific converts
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("cr\t%T\n", (chan<- int)(c))

	// // This doesn't work
	// fmt.Println("-------")
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	// fmt.Printf("cr\t%T\n", (chan int)(cs))

}
