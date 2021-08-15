package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.fname, s.lname, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.fname, p.lname, " - the person speak")
}

type human interface {
	//speak()
	// We can also have an empty Interface.
	// If we have an empty interface then every
	// other type will have no methods atleast :)
	// Other types also have no methods
	// So everything implements the type empty Interface
	// So if we see in Go language something asking for an empty Interface.
	// Then every value can be put in there

	// Go lan doc has
	// func Println(a ...interface{})(n int, err error)
	// Println is asking for unlimited number of type empty Interface.
	// Every value implements type empty interface.
	// "..." means 0 or more values can be put there.
	// Variadic parameters of empty Interface.
	// every type has no methods.
	// They might have one or two methods but they also have no methods as well.
	// So even an int is also of type empty interface.
	// Every value is also of the type empty Interface.
	// Values can be of more than one types.
	// Every type even if they have methods, they also have atleast no methods
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).fname)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).fname)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()

	p1 := person{
		fname: "Dr",
		lname: "No",
	}

	fmt.Println(p1)
	bar(sa1)
	bar(p1)
}
