package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	//p.name = "Miss Money Penny"

	//This is also valid
	(*p).name = "Miss MoneyP"
}

func main() {
	p1 := person{
		name: "James Bond",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
