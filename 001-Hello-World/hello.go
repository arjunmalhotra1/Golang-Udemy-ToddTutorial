package main

import "fmt"

func main() {

	n, e := fmt.Println("Hello Everyone!", 42, true)
	//n, _ := fmt.Println("Hello Everyone!", 42, true)
	//We can choose to ignore the retunr value by mentioning a "_".
	fmt.Println(n)
	fmt.Println(e)
	foo()
	fmt.Println("something more")

	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	bar()
}

func foo() {

	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
