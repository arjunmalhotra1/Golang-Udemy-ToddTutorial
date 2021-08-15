package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// In fan out we have a piece of work. Say we have something that needs to be done over and over again.
// May be say encoding videos. Say, we have a process that runs for encoding videos.
// Say also we have 1000 videos that need to be encoded.
// Say we decide to launch a Go routine for each of the videos that needs to be encoded.
// So we say launch a 1000 go routines and start encoding all the 1000 go routines at once.
// And finally we would want to take the results and bring them back.
// This is Fanning out. Taking a chunk of work and instead of encoding in a sequential fashion,
// encode one and then the next nad then the next, i.e "Serial".
// Fanning out is encoding all of the videos in parallel.
// We can also say instead of Fanning out on all 1000 videos at once,
// We Fan out to only say 10 videos at a time.
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {

	var wg sync.WaitGroup
	for v := range c1 { // Here a Goroutine is launched for all the values in c1. See Throttling Fanout code for more.
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2) // When the Time cosuming work is done the result is dropped into c2.
			// This is Fanning in.
			wg.Done()
		}(v) //We are passing this value "v" into the anonymous function.
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
