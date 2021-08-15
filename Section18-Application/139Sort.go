package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{4, 7, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("----------------")
	sort.Strings(xs)
	fmt.Println(xs)

}
