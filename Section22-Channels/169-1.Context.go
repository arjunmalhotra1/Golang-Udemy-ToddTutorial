package main

import (
	"context"
	"fmt"
)

// concurrent is a tool that we can use with concurrent design pattern, to make sure
// if we have a process which then launches a whole bunch of other goroutines
// when we decide to cancel that one process all of the other launched goroutines are also cancelled.
// We do so, so that we don't leak Goroutines. As it would be using up resources.
// As if we have a process adn part of the process we launch bunch of Goroutines.
// When we close the process we wouldn't want the Goroutines to be still out there. using up the
// resources.
// Read this https://blog.golang.org/context
// https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8
// https://peter.bourgon.org/blog/2016/07/11/context.html

// What we use at Lucid.
// Context can alos help us with passing around of variables whihc are related to a request.
func main() {

	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("---------------")

}
