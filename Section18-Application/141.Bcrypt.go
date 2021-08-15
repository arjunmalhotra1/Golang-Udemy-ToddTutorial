package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "Password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	LoginPassord := "Password1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(LoginPassord))
	if err != nil {
		fmt.Println(err)
		fmt.Println("you can't login")
		return
	}
	fmt.Println("You're logged in")

}
