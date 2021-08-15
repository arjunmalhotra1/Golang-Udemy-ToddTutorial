package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0
	const gs = 100 // 4

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() // Let's other Go routines to run. Similar to saying "Hey pause this one and go run something else"
			v++
			counter = v
			//fmt.Println(counter)
			wg.Done()
		}()
		fmt.Println("Inside for Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count: ", counter)
}
