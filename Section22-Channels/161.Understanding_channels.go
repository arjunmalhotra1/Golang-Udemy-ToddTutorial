package main

import "fmt"

// Channel is a little place where we can send the data and helps us keep the data synchronied.
// Without us having to use mutexes or waitgroups or atomic
func main() {
	// ##### this code dosn't run ######
	c := make(chan int) // It makes a channel
	//<- // This means to take something and put into.
	c <- 42          // This means to put 42 into the channel "c"
	fmt.Println(<-c) // This would then mean to take something off the channel and print it.
	// ##### ######

	//c <- 42 // This is blocked
	// Because when we send and receive on the channel, it is like racers
	// Racing on the trackrace and are passing the baton. they have to pass is hand to hand.
	// The transaction cannot occur until both send and receive can happen at the same time.
	// AT THE SAME TIME.
	// So if they can't happen at the same time. Send&Receive blocks.

	// CHANNELS BLOCK.
	// So our control reaches line 9 and creates a channel.
	// Then on line 11 there's nothing to pull off 42. To pull off the 42.
	// 42 cannot get on the channel until something is ready to pull it off.
	// Pushing and pulling has to happen together.
	// fixed the code in 161Run.Understanding_channels.go

	// From his course notes.
	// channels  block
	// they are like runners in a relay race
	// they are synchronized
	// they have to pass/receive the value at the same time
	// just like runners in a relay race have to pass / receive the baton to each other at the same time
	// one runner can’t pass the baton at one moment
	// and then, later, have the other runner receive the baton
	// the baton is passed/received by the runners at the same time
	// the value is passed/received synchronously; at the same time

}
