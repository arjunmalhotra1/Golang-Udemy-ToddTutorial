package main

import (
	"fmt"
)
//  Interfaces allow us to define behaviour.
// and do Polymorphism.


type secretAgent struct  {
	person
	ltk bool
}


type person struct  {
	fname string
	lname string
}



func (s secretAgent) speak() {
	fmt.Println("I am",s.fname,s.lname," - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am",p.fname,p.lname, " - the person speak")
}

// var x int
// type human interface 
// Notice the similarity on the pattern/structure between the two.
// "keyword" "identifier" "type"

// A value can be of more than one type.
//Like we have sa1, which is of type secretAgent.

// This means any other type that has this method "speak()"
// Also has type human, i.e is also a type human.
type human interface {
	speak()
}

// Todd's Joke - Hey baby if you got this/these method/methods, you are my type.
// Both person and seretAgent have that type

func bar(h human){
	fmt.Println("I was passed into bar",h)
}


func main() {
	sa1 := secretAgent{
		person : person {
			"James",
			"bond",
		},
		ltk : true,
	}

	// sa1 has a value attached to it which is 
	// secret agent, but because it also has a method
	// speak attached to it, it is also of type human.
	// So, sa1 is of type secretAgent and also human.
	// Hence a value can be of more than one type.
	// x := 42
	// This means x is a value of type int.

	fmt.Println(sa1)
	sa1.speak()

	p1 := person{
		fname : "Dr",
		lname : "No",

	}

	fmt.Println(p1)

	// Person has a type speak attached to it.
	// SecretAgent also has a type speak attached to it.
	// And any type that has method speak is also a type human.

	// This is Polymorphism. Poly means "many".
	// "morph" means change.
	// bar can take in many different type, secretAgent
	// as well as person
	// depending on those types we can change
	// what the function does if we wanted to
	// Interfaces allow a value to be of more than one type.
	// Both secretAgent and person have a method speak
	// So any value of type person or any value of type secretAgent
	// They are also of type "Human"
	bar(sa1)
	bar(p1)

}
