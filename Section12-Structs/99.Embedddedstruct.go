package main

import (
	"fmt"
)

// Here we will see how we can take one type and embed it into another type.
type person struct {
	first    string
	lastname string
	age      int
}

type secretAgent struct {
	person
	first string // To demoistrate name colisson.
	// Both secretAgent and person have "fisrt"
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first:    "James",
			lastname: "Bond",
			age:      32,
		},
		first: "Something collison",
		ltk:   true,
	}
	p2 := person{
		first:    "Miss",
		lastname: "Moneypenny",
		age:      27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println("Trying different options")
	//In case of no collison the inner type person is promoted to
	//the outer type secretagent get the person attributes
	// so only sa1.lastname/sa1.age works ust fine
	fmt.Println(sa1.first, sa1.person.first, sa1.lastname, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.lastname, p2.age)
}
