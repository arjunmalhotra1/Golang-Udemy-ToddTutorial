package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //Without this it still ran but exited with an error. because if we never close "c"
		// channel is waiting for someone/something more to come.
	}()

	//receive from the channel
	// This for loop is ranging on the channel "c" until the channel closes.
	// Range polls values from the channel until the channel closes.
	// When the channel closes it pulls out the last value(s) still on the channel.
	// and when there are no more values left on the channel it leaves the range loop.
	// Range loop will contnue to hang around there, until a channel is closed.
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
