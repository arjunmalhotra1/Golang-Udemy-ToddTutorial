package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
	walk()
}

func (p person) walk() {
	fmt.Println("Walking ", p.first)
}

func walkSomewhere(h human) {
	h.walk()
}

func (p *person) speak() {
	fmt.Println("Hello World! ", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}
	saySomething(&p1)
	walkSomewhere(&p1)
	p1.speak()
	p1.walk()
}
