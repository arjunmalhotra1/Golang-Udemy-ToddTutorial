package main

import (
	"fmt"
)


func main() {
	s := "Hello Playground"
	//s = "Hello, Hawaii" //This is not allowed as Strings are immutable
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	for i :=0 ; i<len(s); i++ {
		fmt.Printf("%#U ",s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("")

	for i,v := range s{
		fmt.Printf("at postion %d we have hex %#x\n",i,v)
	}

}
