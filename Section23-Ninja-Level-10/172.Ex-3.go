package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)
	// close(c) I tried closing the channel here but this was wrong.
	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //I had a hard time figuring out where to close the c channel.
	}() // While this go routine is running, "c" has already been passed.
	// channel is closed when the last value has been pulled off.
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	// close(c) I tried closing c here but this was wrong too.

}
