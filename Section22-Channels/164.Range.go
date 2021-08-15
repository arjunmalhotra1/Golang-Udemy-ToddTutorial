package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	//receive
	// Range will keep looping over the channel util the channel is closed.
	// So when we are done with the channel we wnat to close the channel and we do close(c)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to Exit")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		//c <- 42
		c <- i
	}
	close(c)
	// Question. When we close the channel on line 24, are we closing the channel
	// that was on line 20 or the one on line 6?
	// Answer. Try to run this. It is actually closing the primary underlying channel on line 6.
	// To avoid this question he wrote a cleaner code refer 164-A.Range.go file.
}
