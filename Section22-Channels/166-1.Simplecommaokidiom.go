package main

import "fmt"

// Do chekc out https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c

	fmt.Println(v, ok)

	// Now that c is closed this prints 0 and ok is false.
	v, ok = <-c

	fmt.Println(v, ok)
}
