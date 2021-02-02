package main

import (
	"fmt"
)


func main() {
	// switch {
	// case false:
	// 		fmt.Println("This should not print")
	// case (2==4):
	// 		fmt.Println("This should not print")
	// case (3==3):
	// 		fmt.Println("Prints")
	// 		fallthrough // Keyword then it proceeds through
	// case (4==4):
	// 		fmt.Println("Also True, does it print ?")
	// 		fallthrough
	// case (7 == 9):
	// 		fmt.Println("Not true1")
	// 		fallthrough
	// case (11 == 14):
	// 		fmt.Println("Not True 2")
	// 		fallthrough
	// case (15==15):
	// 		fmt.Println("True 15")
	// 		fallthrough
	// default:
	// 	fmt.Println("This is default")

	///////////////////////////////////////////////////////////////	

	// switch {
	// case false:
	// 		fmt.Println("This should not print")
	// case (2==4):
	// 		fmt.Println("This should not print")
	// default:
	// 	fmt.Println("This is default")
	// }

	///////////////////////////////////////////////////////////////////

	// switch "Bond"{
	// case "MoneyPenny":
	// 		fmt.Println("Miss Money")
	// case "Bond":
	// 		fmt.Println("Bond James")
	// case "Q":
	// 		fmt.Println("This is q")
	// default:
	// 	fmt.Println("This is default")
	// }

	/////////////////////////////////////////////////////////////
	n:="Bond"
	switch n{
	case "MoneyPenny", "Bond", "Dr No":
			fmt.Println("Miss Money, Bond or Dr no")
	case "M":
			fmt.Println("This is M")
	case "Q":
			fmt.Println("This is q")
	default:
		fmt.Println("This is default")
	}



}