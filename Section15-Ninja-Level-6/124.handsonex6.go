package main

import "fmt"

func main() {
	// func() {
	// 	fmt.Println("Hello World")
	// }()

	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

}
