package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var gs int = 100
var counter int = 0
var mu sync.Mutex

func main() {
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			// mu.Unlock() // IMP - Here I was still getting one Data Race.
			// This is because the Go routines were still competing for the counter to be printed.
			// Hence moving the Unlock after the "fmt.Println(counter)" fixed it.
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("end counter", counter)
}
