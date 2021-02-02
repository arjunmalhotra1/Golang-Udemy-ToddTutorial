package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James" : 32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["barnabas"])

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	// To add a key we simply do
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Now switching to the slices")
	xi := []int{4,5,7,8,9,42}

	for i, v := range xi {
		fmt.Println(i,v)
	}

}
