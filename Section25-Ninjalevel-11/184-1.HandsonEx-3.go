package main

import "fmt"

type customErr struct {
	info string
}

func (cerr customErr) Error() string {
	return fmt.Sprintf("Here is the err %v", cerr.info)
}
func main() {
	err1 := customErr{
		info: "Hello Error, need more Coffee",
	}
	foo(err1)
}

func foo(err error) {
	fmt.Println("foo ran - ", err) // This works and this is what was needed.

	//fmt.Println("foo ran - ", err, "\n", err.info) // This does not work.
	//err.info undefined (type error has no field or method info)
	// Because info is not attached to error but info is attached to customErr
	// To use info we will haev to use Assertion.
	// Assertion is different from conversion.

	fmt.Println("foo ran - ", err, "\n", err.(customErr).info)
	// We are asserting that "err" is of type customErr.
	// This is static programming.
	// But we have to explicitly tell the run time compiler that this is customErr and let's pull info field
	// off of it and that's going to run.

}
