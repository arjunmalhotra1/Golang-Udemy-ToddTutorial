package main

import (
	"fmt"
)

// Struct is an aggregarte Data Structure that allows us to compose together 
//  values of different tyeps, hence it is an aggregate data structure.
// Or we can say we are composing the different values.
// aka complex data types.

type person struct{
	first string
	lastname string
	age int
}

func main() {
	p1 := person{
		first: "James",
		lastname: "Bond",
		age: 32,
	}

	p2 := person {
		first: "Miss",
		lastname: "Moneypenny",
		age: 27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("Other printing options")
	
	fmt.Println(p1.first,p1.lastname,p1.age)
	fmt.Println(p2.first,p2.lastname,p2.age)

}
