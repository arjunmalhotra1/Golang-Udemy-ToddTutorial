package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// Send to the Channel
func send(even, odd chan<- int, quit chan<- bool) {
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
func receive(odd, even <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("The value received from the even channel: ", v)
		case v := <-odd:
			fmt.Println("The value received from the odd channel: ", v)
		case i, ok := <-quit: // When quit is closed sends over false and false. i and ok are both false.
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
