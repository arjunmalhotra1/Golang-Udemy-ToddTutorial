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
	speak()
}

//Assertion

// We cannot jsut do h.fname
// As this might be just secretAgent or person
// This is polymorphism as either peson or secretAgent
// can be passed into bar as a human
func bar(h human) {
	//fmt.Println("I was passed into bar",h) // This Works
	//fmt.Println("I was passed into bar",h.fname) // This doesn't Work
	// We can make it work using a switch statement

	// h.(person) this means we are asserting, that this "h" is type person.
	// h.(secretAgent) this means we are asserting, that this "h" is type secretAgent.
	// we know that because we call h.(peron) and h.(secretAgent) inside switch.
	// Now we can call "h.(person).fname" and "h.(secretAgent).fname"
	// "h.(person).fname" says I know this "h" is type person as we just asserted it
	// So now let's just call "fname" of the type person.
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).fname)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).fname)
	}
}

type hotdog int // Tihs is for the reference of conversion.

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

	// Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) // Converted type hotdog into an int.
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
