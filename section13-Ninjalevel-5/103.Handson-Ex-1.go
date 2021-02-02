package main

import (
	"fmt"
)

type person struct{
	fname string
	lname string
	favicecream []string
}

func main() {
	p1 := person{
		fname : "James",
		lname : "Bond",
		favicecream: []string {
			"Chocolate",
			"Vanilla",
			"martini",
		},
	}
	p2 := person{
		fname: "Miss",
		lname: "Monepenny",
		favicecream: []string{
			"hazelnut",
			"strawberry",
		},
	}
	//fmt.Println(p1)
	//fmt.Println(p2)

	fmt.Println(p1.fname)
	fmt.Println(p1.lname)

	for i, v:= range p1.favicecream{
		fmt.Println(i,v)
	}

	fmt.Println(p2.fname)
	fmt.Println(p2.lname)

	for i, v := range p2.favicecream{
		fmt.Println(i,v)
	}



}
