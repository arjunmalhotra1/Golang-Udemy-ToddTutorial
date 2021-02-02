package main

import "fmt"

var y = 42 // We have declared variable y.
			//It can only store the values of type int.

// WE are declaring a VARIABLE iwth IDENTIFIER "z"
// is of TYPE string
// and I assign the VALUE "Shaken not stirred"
var z = `James said, "Shaken not Stirred"`
// ` ` back ticks will print everything between then even the quotes and newline 
// var z = `James said, 


//"Shaken not Stirred"`

// GO Lang is a Static programming language.
// and not Dynamic programming language.
// a variable is declared to hold a VALUE of a certain TYPE.
// var z string = "Shaken not Stirred" //We can do this as well. 
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n",y)
	//z = 43 // This gives error as Golang is Static typed Programming langage, unline JS which is dynamic.
	fmt.Println(z)
	fmt.Printf("%T\n",z)

}
