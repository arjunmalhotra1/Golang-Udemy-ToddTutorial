package main

import (
	"fmt"
)

// If we want to keep our code neat and clean
// Do not have extraneous types where we don't need them

func main() {
	p1 := struct {
		first  string
		second string
		age    int
	}{
		first:  "James",
		second: "Bond",
		age:    32,
	}
	fmt.Println(p1)
}

// WE create VALUES of a certian TYPE that are stored in VARIABLES.
// and those VARIABLES have IDENTIFIERS.
