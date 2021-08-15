package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var gs int = 100
var counter int = 0

func main() {
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end counter", counter)
}
