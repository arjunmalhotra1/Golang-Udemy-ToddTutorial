package main

import (
	"fmt"
)
 const (
	 a = iota // a is initilaized to 0 
	 b 		 // b is now 1
	 c 		 // c is now 2
 )
 const (
	 d = iota // d is initalized to be 0
	 e		  // e is 1
	 f		  // f i s2
 )

const (
	a = iota // a is initilaized to 0 
	b 		 // b is now 1
	c 		 // c is now 2
	d = iota // d is 3
	e		  // e is 4
	f		  // f is 5
)


func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c )
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
