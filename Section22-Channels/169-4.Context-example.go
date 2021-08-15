package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1: ", ctx.Err())
	fmt.Println("num gortins 1: ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select { // context document says that Done is used inside the select statement.
			//The Done method returns a channel that acts as a cancelation signal to functions running on behalf of the
			//Context: when the channel is closed, the functions should abandon their work and return.
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2: ", ctx.Err())
	fmt.Println("num gortins 2: ", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("Cancelled  context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3: ", ctx.Err())
	fmt.Println("num gottins 3: ", runtime.NumGoroutine())
}
