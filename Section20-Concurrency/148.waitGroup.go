package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

//Func init runs before the main for any initializations and we can have as many func inits as we want.
// func init(){}
func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo() // If we don't have ws.Add(1) then the main go routine doesn't wait.
	// During the time taken to spin up a goroutine, bar and everything beyond it ran.
	// And func main() shut down.
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // Here we are waiting for all the things we added in the line 19.
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar ", i)
	}
}
