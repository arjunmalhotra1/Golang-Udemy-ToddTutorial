package main

import (
	"fmt"
)


func main() {
	mp := map[string]int {
		"James" : 32,
		"Moneypenny": 27,
	}

	fmt.Println(mp)

	delete(mp, "James")
	fmt.Println(mp)

	// We can delete soemthing that never existed in the map as well.
	delete(mp, "Ian fleming")
	fmt.Println(mp)

	fmt.Println(mp["Moneypenny"])
	fmt.Println(mp["Ian Fleming"])

	// So to delete something forom the map we us
	// comma ok idiomatic way in golang.
	if v, ok := mp["Moneypenny"]; ok {
		fmt.Println("value:",v)
		delete(mp, "Moneypenny")
	}
	fmt.Println(mp)

}
