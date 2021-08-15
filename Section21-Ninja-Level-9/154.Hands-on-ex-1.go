package main

import (
	"fmt"
	"runtime"
	"sync"
)

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(1)
// 	go func1()
// 	fucn2()

// 	wg.Wait()

// }

// func func1() {
// 	fmt.Println("Go routine 1")
// 	wg.Done()
// }
// func fucn2() {
// 	fmt.Println("Go routine 2")
// }

func main() {

	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Hello from routine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from routine 2")
		wg.Done()
	}()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End gs", runtime.NumGoroutine())
}
