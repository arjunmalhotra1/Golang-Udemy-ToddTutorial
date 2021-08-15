package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			//close(c) // This is not going to work because we would want to to close c not 10 times but only once.
		}()
	}

	// for x := range c {
	// 	fmt.Println(x)
	// }
	// close(c)// This resulted in a deadlock too

	// So he did this, instead of using range we know how many values are going to be put onto that channel.
	// So we do this:
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("About to exit")
}
