// This solution is using the wait group. A student had it in the comments of this lecture.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			for j := 0; j < 10; j++ {
				c <- j
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	for x := range c {
		fmt.Println(x)
	}
}
