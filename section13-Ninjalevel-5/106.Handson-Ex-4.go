package main

import (
	"fmt"
)


func main() {
	s := struct{
		first string
		friends map[string]int
		favdrinks [] string
	}{
		first: "James",
		friends: map[string]int {
			"Todd" : 132465,
			"Bob" : 46456,
		},
		favdrinks : []string {
			"martini",
			"rum and coke",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favdrinks)

	for k, v := range s.friends {
		fmt.Println(k,v)
	}

	for i, val := range s.favdrinks {
		fmt.Println(i,val)
	}
}
