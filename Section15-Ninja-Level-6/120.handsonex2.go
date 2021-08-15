package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	s := foo(xi...)
	fmt.Println("foo sum is ", s)
	t := bar(xi)
	fmt.Println("bar sum is ", t)
}

func foo(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

func bar(sliceOfInt []int) int {
	sum := 0
	for _, v := range sliceOfInt {
		sum += v
	}
	return sum
}
