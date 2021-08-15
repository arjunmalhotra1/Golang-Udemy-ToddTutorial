package main

import (
	"fmt"
	"sync"
)

// Fan in and Fan out are two concurrency patterns in Golang.
// Say we have a chunck of work and we don't know how much work it is going to be.
// So we fan it out to as many Go routines as possible.
// They all will be working on that work as much as possible.
// When they get results we will Fan all the results in to another channel and we
// will get the channel with just the results.
func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
// We are passing bi-directional channels odd and even to be caught as a send only channel
func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}
	close(even)
	close(odd)
}

//receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
