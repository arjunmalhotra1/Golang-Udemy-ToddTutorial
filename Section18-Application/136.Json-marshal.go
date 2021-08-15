package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		"James",
		"Bond",
		32,
	}

	p2 := person{
		"Miss",
		"MoneyPenny",
		27,
	}

	people := []person{
		p1,
		p2,
	}
	fmt.Println(people)

	//Say we need to send this data to somebody and so we will need to Marshall it.
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) // This returned nothing when the fields of the struct
	// had the first letter as lower case.
	// Even though the "struct" name "person" has lower first letter but the inner fields have
	// the first letter as upper case so that the Json will be able to export them.
}
