package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("Defer inside foo")
	}()
	fmt.Println("Inside foo")
}

func bar() {
	fmt.Println("Inside bar")
}
