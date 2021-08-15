package main

import "fmt"

func main() {
	//c := make(chan int) // this one failed.
	// We can have buffered channel. Buffered channel is a channel that will allow a certain values
	// to sit in regardless of if a certain value is ready to be pulled off or not.
	c := make(chan int, 1) // This means that this channel will allow one value to sit in.
	c <- 42                // 42 sits in the code so we have room to store one value.

	fmt.Println(<-c)
}

// There are unsuccessfull buffered channels as well.
// See 161UnsuccRun3.Understanding_channels.go
