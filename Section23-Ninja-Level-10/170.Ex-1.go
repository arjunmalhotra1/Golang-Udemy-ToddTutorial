package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//Method - 1
	//c := make(chan int, 1) // Buffered channel

	// Method - 2
	// go func() {
	// 	c <- 42
	// }()

	fmt.Println(<-c)

}
