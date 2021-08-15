package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 42
	c <- 43

	fmt.Println(<-c) // Prints 42
	//fmt.Println(<-c) // Prints 43 if we have | c := make(chan int, 2) |
	// This code is fine till line 7.
	// Beyond that on line 8 it gets blocked.
	// Saying "oh! we only have room for one value"
	// We can do make(chan int, 2)
	// But if we keep increasing the size we will never know when the buffer is full.

	// Ignore the reasoning if you can't understand it.
	// Bill Kennedy's advice - Try to stay away from Buffered channels.
	// Because of the unpredicable nature. So try to use UnBuffered channels.
	// As we would want to make sure that our code is build in such a way that ther eis always an
	// interlocking way such that a pass of the baton happens.
	// That is exchange of the baton between the putting on the channel and taking off the channel.

	// Int he langugae specification.
	// DO NOT COMMUNICATE BY SHARING MEMORY INSTEAD SHARE MEMORY BY COMMUNICATING

}
