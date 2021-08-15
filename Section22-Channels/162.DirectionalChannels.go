package main

import (
	"fmt"	"hash/crc32"
)


func main() {
	c := make(chan int, 2) // We are able to both send and receive on this channel.
	c <- 42                // Send into the channel
	c <- 43                // Send into the channel

	fmt.Println(<-c) // Receive forom the channel
	fmt.Println(<-c) // Receive forom the channel

	// We can pass channels off which we only read as the values will be put in
	// at different place.

	fmt.Println("---------------")
	fmt.Printf("%T\n", c)

	// READ LEFT TO RIGHT
	//chan<- int SEND TO THE CHANNEL AN INT
	d := make(chan<- int, 2) // Here the channel is only send channel.
	// you can only send to the channel and not read/receive of it.

	d <- 42
	d <- 43

	//fmt.Println(<-d) // This gives us an error as we are trying to read form the send only channel.

	// READ LEFT TO RIGHT "<-chan int"
	// We are receiving ofrom channel int
	e := make(<-chan int, 2) //This is a receive only channel.
	// You can only read/receive from it. But not send to it.

	//e <- 42 // This gives an erros as we are trying to send to the read only channel

	// SPECIFIC TO GENERAL DOESN'T WORK
	// c:= make(chan int)
	// cr:= make(<-chan int)
	// cs:= make(chan<- int)
	// SPECIFIC TO GENERAL DOESN'T WORK.
	// c=cr // "cr" is specific(receive channel) and "c" is general(bidirectional). 
	// Similarly when we go from 46.34(float) to an int more general
	// Then we lose data.
	// c=cs

	// THIS WORKS
	// GENERAL TO SPECIFIC WORKS
	// cr = c // c is general (bidirectional) and cr is specific (receive)
	// cs = c
}
