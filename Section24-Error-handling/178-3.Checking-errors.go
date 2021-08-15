package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// In this example we create a file.
func main() {
	f, err := os.Create("names.txt")
	if err != nil { // That is if there is an error we will print an error.
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond")

	io.Copy(f, r) // So it reads from r and writes to f
}
