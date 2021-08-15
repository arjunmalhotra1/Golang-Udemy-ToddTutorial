package main

import "fmt"

func main() {
	c := make(chan int)

	// Launching the two go routines.
	//send
	go foo(c)

	//receive
	// "go bar(c)" he had this earlienr then he removed the "go" keyword.
	bar(c) // The recieve channel in bar is blocked until the funciton "foo" writes something to the
	// channel
	//Refer 161run.Understanding_channels.go

	// When I did
	// foo(c) &
	// go bar(c)
	// It threw an error. As in 161.Understanding_channel.go
	// There is nothing to pull off 42 as bar ran off as a parallel excutable go routine.
	// So Channel gets blocked.

	// But When I did
	// go foo(c) &
	// go bar(c)
	// Both "go routines" go off to run and the main doesn't wait and the program prints
	// About to exit and terminates"

	fmt.Println("About to Exit")
}

//send
func foo(c chan<- int) {
	c <- 42

}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
