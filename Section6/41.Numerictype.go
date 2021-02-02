package main

import (
	"fmt"
)

var x int
var x1 int8 = -128 // -129 does not work here.
var y float64

func main() {
	x = 42
	y = 42.345
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
	fmt.Println(x1)
	fmt.Printf("%T\n",x1)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	
}
