package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", sum(2, 3))
	fmt.Println("4 + 5 = ", sum(4, 5))
	fmt.Println("6 + 8 = ", sum(6, 8))
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
