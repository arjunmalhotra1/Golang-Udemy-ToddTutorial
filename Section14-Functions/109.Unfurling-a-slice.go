package main

import (
	"fmt"
)


func main() {
	//xi := [] int {2,3,4,5,6,7,8,9} // It is the same underline array 
									 // that is passed tot he function sum(xi ...)
	//x := sum(xi...)
	x := sum("james")
	fmt.Println("The total is", x)


}
// Say we were to pass String and variadic parameters
// We cannot do 
//func sum(x ..int, s string)
//But 
//We can do
//func sum(s string, x ...int)
//So variadic parameter has to be the final parameter passed in the function
//Variadic "...""  means 0 or more so we could pass nothing as well.
//func sum(x ...int) int {
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // underlying array also didn't get create if 
						//no arguments are passed.

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for the item in position",i,"We are adding",v,"to the total which is now",sum)
	}
	return sum
}