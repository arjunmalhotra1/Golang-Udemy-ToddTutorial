package main

import (
	"fmt"
)
type hotdog int
var x hotdog

func main() {
	//a:=42
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
}
