package main

import (
	"fmt"
)

func main() {
	//m := map [string]int{} // This is a composite literal.
	// [string]int is the entire type of the map.

	m := map [string]int{
			"James":32,
			"Miss MoneyPenny":27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) // This returns 0.
	// If we enter a key and that key does not have a corresposnding value
	// stored int he map that key will return a 0 value.
	// But his can be confusing, if say we were storing marks for an exam.

	//Identifier to store the check is often OK
	v, ok := m["Barnabas"]

	fmt.Println(v)
	fmt.Println(ok) // returns false.

	if v, ok := m["Barnabas"]; ok{
		fmt.Println("This is the if Print",v)
	}

	if v, ok := m["Miss MoneyPenny"]; ok{
		fmt.Println("This is the if Print for Miss Moneypenny",v)
	}





}
