package main

import "fmt"

func main() {
	c := make(chan int)

	// Now this Go Routine is launched off and is running.
	// We have now 2 Goroutines running in parallel in our concurrent design pattern.
	go func() {
		c <- 42
	}()
	// Again line 11 blocks.
	// But because this Go Routine has been sent off to run on it's own.
	// The "main" Go Routine after firing off the go routine on line 10, just continues.
	fmt.Println(<-c)
	// At line 16, line 11 is ready to put it and is blocking so far.
	// Now At line 16 it is ready to take the 42 off the channel.
	// Both are runing at the same time and pass the baton.
	// On line 11 is ready to put the value in the channel and then
	// On line 16 it is ready to take the value off the channel.
	// This is how we have program coordination with concurrent design.

	// The second way to do this is in 161RUN2.Understanding_channels.go
}
