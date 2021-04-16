package main

import (
	"fmt"
)

func main(){
	counter := 0

	const gs = 100

	for i:=0; i < gs; i++{
		go func(){
			v := counter
			v++
			counter = v
		}()
	}
	fmt.Println("Hello, playground")
}