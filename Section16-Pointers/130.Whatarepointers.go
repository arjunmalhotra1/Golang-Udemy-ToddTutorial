package main

import (
	"fmt"
)


func main() {
	x:=42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n", &x)

	//We cannot do this
	//var b int = &x // The compiler says 
	// Cannot use &x which is pointer to an int as type int in assignment
	// We can't assign a pointer to an int.
	// We can assign a pointer to only a pointer.

	// This is good.
	//var b *int = &x

	//This is good as well
	b := &x

	fmt.Println(b)
	fmt.Println(*b) // This * is an operator
	// We are dereferencing an address here.

	// Had x been a char
	// b would have been *char
	// Whihc would have meant it is pointing to a `char`

	fmt.Println("Adrerss of x ",&x)
	fmt.Println("Value of x ",*&x)

	// & -> gives you the address and,
	// * -> gives you the value stored at an address when you have the address

	*b = 43
	fmt.Println("Updated  Value of x",x)

	// *int
	// This * is part of a type.
	// It is a pointer to where int is stored.
}
 
