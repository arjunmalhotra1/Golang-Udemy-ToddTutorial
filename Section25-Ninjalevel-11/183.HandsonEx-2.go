package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		// fmt.Println(err)
		// return
		//log.Println(err)
		log.Fatalln(err) // if we do this then we don't need the return because the fatal calls os.Exit(1)

	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("There was an error in toJson: %v", err) // This works - 1st methd=od
		//return []byte{}, fmt.Println("There was an error in toJson: %v", err) // This doesn't work.
		//return []byte{}, errors.New("There was an error in toJson: %v", err) // This doesn't work.
		//As we give 2 arguments to the errors.New but it takes only one string.
		return []byte{}, errors.New(fmt.Sprintf("There was an error in toJson: %v", err))

	}
	return bs, nil

}
