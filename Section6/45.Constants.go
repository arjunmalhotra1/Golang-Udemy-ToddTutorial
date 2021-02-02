package main

import (
	"fmt"
)

//const a = 42
//const b = 42.78
//const c ="Jams Bond"
//Untyped constants
// const (
// 	a = 42
// 	b = 42.78
// 	c ="James bond"
// )

const (
	a int = 42
	b float64 = 42.78
	c string ="James bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

