package main

import (
	"fmt"
)


func main() {
	x := map[string][]string {
		"James" : []string{"Shaken, no Stirred", "Martinis", "Women"},
		"moneypenny_miss" : []string{"James, Bond", "Literature", "Computer Science"},
		"no_dr" : []string{"Being evil", "Ice Cream", "Sunsets"},
	}
	fmt.Println(x)

	for k,v := range x{
		fmt.Println("This record for",k)
		for k1,v1 := range v{
			fmt.Println("Interests \t",k1,v1)
		}
	}
}
