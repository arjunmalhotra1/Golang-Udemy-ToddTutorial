package main

import (
	"fmt"
)


func main() {
	
	y := []string{"James","Bond","Shaken, not stirred"}
	z := []string{"Miss","MoneyPenny","Helloooo,James"}
	x := [][]string {z, y}
	fmt.Println(x)

	for i, xs := range x {
		fmt.Println("record: ",i)
		for j, val := range xs{
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}
