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

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area ", s.area())
}

func main() {
	c := circle{5}
	// info(c) // This gives an error because the function area is associated with
	// simply "c".

	fmt.Println(c.area()) // This code runs.
	// Method set determines the INTERFACES that the type implements.
	// Refer
	// https://tour.golang.org/methods/6
	// then
	// https://tour.golang.org/methods/7
	// comments in this lecture
	// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922318#questions/5355558

}
