package main

import (
	"fmt"
)

type person struct{
	fname string
	lname string
	ice []string
}
func main() {
	p1 := person {
		fname:"James",
		lname:"bond",
		ice:[]string{
			"martini",
			"scotch",
			"rum",
		},
	}
	p2 := person{
		fname: "Miss",
		lname:"Moneypenny",
		ice: []string{
			"chocolate",
			"strawberry",
			"vannila",
		},
	}

	m := map[string]person{
		p1.lname : p1,
		p2.lname : p2,
	}
	//fmt.Println(m["bond"])
	//fmt.Println(m["Moneypenny"])
	//Other way to print this is 
	for k,v := range m{
		fmt.Println(k)
		//fmt.Println(v)
		fmt.Println(v.fname)
		fmt.Println(v.lname)
		for i,val := range v.ice {
			fmt.Println(i,val)
		}
		fmt.Println("-----------------")
	}

}

