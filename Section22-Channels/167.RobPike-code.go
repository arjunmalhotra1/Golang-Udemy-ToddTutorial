package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring; I'm leaving")

}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ { // This loop will keep looping forever
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string) // "c" on line 20 and this c on line 32 have totally different scopes and
	// hence are totally different chanels
	go func() {
		for {
			c <- <-input1 // Take value off of input1 and put it on to channel c
		}
	}()

	go func() {
		for {
			c <- <-input2 // Take value off of input2 and put it on to channel c
		}
	}()
	return c // returns a receive only channel. As the return type is specific, which is receive only.
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
