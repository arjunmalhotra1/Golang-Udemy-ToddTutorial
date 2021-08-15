package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is caled, defrred functions don't run")
}
