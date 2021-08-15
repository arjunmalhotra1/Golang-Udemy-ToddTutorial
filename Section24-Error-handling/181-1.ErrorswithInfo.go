package main

import (
	"errors"
	"log"
)

// When an error is thrown, we can have information about
// where it occurred and what happened
// so that we can adda s much information as we want.
// We can add information to our errors by using:
// errors.New()

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, errors.New("norgate Math : square root of negative number")
	}
	return 42, nil
}
