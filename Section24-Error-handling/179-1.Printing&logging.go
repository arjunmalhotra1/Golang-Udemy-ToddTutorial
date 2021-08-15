// With log.Fatalln() - Deferred fucntions do not run
// With log.Panicln() - Deferred fucntions run

package main

import (
	"fmt"
	"log"
	"os"
)

//Fataln is equivalent to Println followed by call to os.Exit(1).

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err) // Prints - err happened open no-file.txt:
		// The system cannot find the file specified.

		// We get date time stamp with the log.
		log.Println("err happened", err) // Prints - 2021/06/03 10:12:03 err happened open no-file.txt:
		// The system cannot find the file specified.

		// Exit causes the current program to exit with the given status code.
		//Conventionally, code zero indicates success, non-zero an error.
		// The program terminates immediately; deferred functions are not run.
		log.Fatalln(err) //Fatal function calls Os.Exit(1) function after writing the log messages.
		// Fatal function is equivalent to Println() followed by a call to os.Exit(1)

		log.Panicln(err)
		// Deferred functions are run by panic
		//Panic built-in fucntion stops the normal execcution of the current goroutine.
		// Panicln is equivalent to Println() followed by a call to panic()
		// Panic printed this:
		// panic: open no-file.txt: The system cannot find the file specified.
		//goroutine 1 [running]:
		//main.main()
		//C:/Users/Arjun/Documents/Udemy/Go/Section24-Error-handling/179-1.Printing&logging.go:28 +0x7b
		//exit status 2
		// Read the below link
		//https://golang.org/pkg/builtin/#panic

		panic(err)
		// Panic fucntions call panic after writing the log message.
		// Similar to calling log.Panicln()
	}

}
