package main

import "fmt"

// Channels block and
// We can have a range clause which alos blocks until the channel is closed.

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send into the channel
	go send(eve, odd, quit)

	//receive from the channel
	receive(eve, odd, quit)

	fmt.Println("About to exit")

}

//func send(evecs chan<-int,oddcs chan<-int, quit)

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e) // He later removed these 2 close statements as they were printing extra 0s off even and odd channel.
	// close(o)
	q <- 0

}

func receive(e, o, q <-chan int) {
	//Infinite for loop
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel: ", v)
		case v := <-o:
			fmt.Println("from the odd channel: ", v)
		case v := <-q:
			fmt.Println("From the quit channel: ", v)
			//close(q) // This gave an error, I don't know why. The error says
			// "invalid operation: close(q) (cannot close receive-only channel)"
			return // As we would want to return from our function at some point.
		}
	}
	// Select statemnt is going to pull the values off of different channels.
	// Select statement is going ot say "Hey! which of these channel I am going to pull the value off of?"
	// Select will keep taking the value off the channel until it takes the value off the quit channel.
	// Then return.
}
