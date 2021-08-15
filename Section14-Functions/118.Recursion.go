package main

import "fmt"

func main() {

	fmt.Println(4 * 3 * 2 * 1)

	n := factorial(4)
	fmt.Println("RecursionFactorial product", n)

	n2 := factorialLoop(4)
	fmt.Println("LoopFactorial product ", n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func factorialLoop(n int) int {

	//var prod int = 1
	// for _, v := range n {
	// 	prod = prod * v
	// }

	prod := 1
	for ; n > 0; n-- { // Since we needed only the condition and the post we left the init empty.
		prod *= n
	}
	return prod
}
