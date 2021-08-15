package main

import "fmt"

func main() {
	n := foo()
	fmt.Println(n)

	x, s := bar()
	fmt.Print("x ", x, " s ", s)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big brother is watching"
}
