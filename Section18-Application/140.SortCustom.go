package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type ByAge []person

// By Age implements the sort.Interface for []Person based on
// Age field.
// https://pkg.go.dev/sort#Interface
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i int, j int)  { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i int, j int)  { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].first < bn[j].first }

func main() {
	p1 := person{"James", 32}
	p2 := person{"MoneyPenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	// Since people is of type []person (slice person) &
	// ByAge is of type slice person.
	// We are nothing but doing type conversion here.
	// We are trying to convert people of to type of ByAge.
	// and since it is of type ByAge it has the methods Len, Swap and Less attached to it.
	sort.Sort(ByAge(people))
	// When we do type ocnversion, ByAge is implicitly implementing the interface from package sort.
	// Because https://pkg.go.dev/sort#Sort "sort.Sort" asked for the interface.
	fmt.Println(people)

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
