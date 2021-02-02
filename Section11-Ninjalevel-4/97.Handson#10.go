package main

import (
	"fmt"
)


func main() {
	x := map[string][]string{
		"James":[]string{"Shaken, Not Stirred","Martinis","Women"},
		"monepyeeny_miss":[]string{"JamesBond","Literature","ComputerScience"},
		"no_dr": []string{"Being_Evil","IceCream","sunset"},
	}

	x["fleming"] = []string{"steaks","cigars","espionage"}

	for k,v := range x{
		fmt.Println("This record is for",k)
		for i,v1 := range v{
			fmt.Println("\t",i,v1)
		}
	}
	
	fmt.Println("\n Deleting \n")

	delete(x,`no_dr`)
	for k,v := range x{
		fmt.Println("This record is for",k)
		for i,v1 := range v{
			fmt.Println("\t",i,v1)
		}
	}
}
