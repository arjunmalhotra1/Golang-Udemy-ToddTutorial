package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(s shape) {
	//x := s.area()
	//fmt.Println(x)
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 12.345,
	}
	s := square{
		side: 15,
	}

	info(c)
	info(s)

}
