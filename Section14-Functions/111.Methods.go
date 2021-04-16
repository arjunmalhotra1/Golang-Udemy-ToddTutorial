package main

import (
	"fmt"
)
type person struct{
	fname string
	lname string
}

type secretAgent struct{
	person
	ltk bool
}
// func (r receiver) identifier(parameters) (return(s)) { code  }

// We are attaching the function speak() to anybody
// Who has a type secretAgent
// Any value of the type secretAgent has access to the function speak 
// When we have a reciever it attaches the function to the receiver type.
// Here secret Agent which happens to be a struct.

func (s secretAgent) speak() {
	fmt.Println("i am", s.fname, s.lname)
}


func main() {
	sa1 := secretAgent {
		person : person {
			"James",
			"Bond",
		},
		ltk : true,
	}

	sa2 := secretAgent {
		person : person {
			"Miss",
			"Moneypenny",
		},
		ltk : false,
	}

	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()
}
