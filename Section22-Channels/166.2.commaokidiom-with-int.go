package main

import "fmt"

// In 165 we had a few zeroes and when we closed the channels. That is why he said that we'll look into
// comma ok idiom.
// If the channel is closed you can not write to channel but you can read from closed channel until emptyness.
// You could not close channel only if you do not read from it at all.

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// Send to the Channel
func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

//Receive from the channel
func receive(odd, even, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("The value received from the even channel: ", v)
		case v := <-odd:
			fmt.Println("The value received from the odd channel: ", v)
		case i, ok := <-quit: // When quit is closed it will pass 0 and false.
			// We use comma ok idiom to chekc if the channel is closed.
			if !ok {
				fmt.Println("from comma ok  ", i, ok)
				return
			} else {
				fmt.Println("from comma ok  ", i)
			}
		}
	}
}
