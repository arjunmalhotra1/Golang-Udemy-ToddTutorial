package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 { // Receiver (c circle) is a non-pointer receiver
	return math.Pi * c.radius * c.radius
}

// When it is like this then only "info(&c)" will work. "info(c)" wil give us an error.
// func (c *circle) areaPointer() float64 {
// 	return math.Pi * c.radius * c.radius
// }

func info(s shape) {
	fmt.Println("area: ", s.area())
}
func main() {

	c := circle{5}
	info(c)  // This is example of non-pointer receiver and non-pointer value.
	info(&c) // This is example of non-pointer receiver and pointer value.

	fmt.Printf("%T\n", &c) // This prints "*main.circle"
	// which means &c is a pointer to type circle from package main.

	// The below lines are from - Video 149 Method sets Revisited - Section 20 - Concurrency
	// go to 149.Methodsets-Revisited.go in Section 20 -Concurrency.
	//fmt.Println(c.area())
}
