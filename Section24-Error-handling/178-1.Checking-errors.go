package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello")

	if err != nil { // err==nil means there is no error and err!=nil means there is an error.
		fmt.Println(err)
	}
	fmt.Println(n) // Although "Hello" is 5 characters this prints 6 is beacause of the new line character.
	// if we do
	// n,err := fmt.Print("Hello") then fmt.Println(n) prints 5.

}
