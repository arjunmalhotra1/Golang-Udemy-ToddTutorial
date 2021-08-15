package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		// Without the close(c). I got the deadlock error
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit.")
}
