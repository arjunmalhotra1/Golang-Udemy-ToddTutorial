package main

import "fmt"

// Callback is a function passing another fucntion as an argument.
// As we saw that functions are first calss citizens in golang.
// Call back - Passing a function in as an argumet to another function.

func main() {
	ii := []int{1, 2, 3, 4, 5}
	s := sum(ii...)
	fmt.Println("sum of all numbers ", s)

	s2 := even(sum, ii...)
	fmt.Println("sum of even numbers ", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Sum of odd numbers", s3)

}

func sum(xi ...int) int {
	//fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// Here is a call back example.
// We are passing in this funciton.
// f func(xi ...int) int
// We are passing in a function that takes in a variadic parameter of ints assigned to
// the "xi". That returns an int.
// We are passing in that function as an argument. Which is assigned to f.
// In addition we are also passing in variadic parameter into "vi"
func even(f func(xi ...int) int, vi ...int) int {

	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}

func odd(f func(xi ...int) int, vi ...int) int {

	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}
