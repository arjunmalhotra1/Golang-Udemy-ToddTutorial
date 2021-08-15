package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u user) testNoPointerReceiver() {
	fmt.Println("No Pointer Receiver --> success")
}

func (u *user) testWithPointerReceiver() {
	fmt.Println("With Pointer Receiver --> success")
}

// implementing the notifier interface method notify(). Using a pointer receiver of user
func (u *user) notify() {
	fmt.Println("Sending message to", u.name, " --> success")
}

func main() {
	u := user{"John Doe"}

	fmt.Println("In (1) (2) and (3) ==> you have to refer to table that shows method sets from the method sets perspective")

	// (1) Calling regular non interface methods using user variable ie addressable _ with and without pointer receiver
	fmt.Println("\n(1) [Non-interface methods] Calling regular non interface methods using user variable ie addressable _ with and without pointer receiver")

	fmt.Println("--> Will run method using user value on a NON pointer receiver method")
	u.testNoPointerReceiver()
	// works

	fmt.Println()

	fmt.Println("--> Will run method using user value on a pointer receiver method")

	u.testWithPointerReceiver()
	// !!works: u is a variable ie addressable. Compiler will implicitely put it as (&u).testWithPointerReceiver() so it will work with function's pointer receiver

	fmt.Println()

	// (2) Calling regular non interface methods using object literal ie non addressable _ with and without pointer receiver
	fmt.Println("(2)  [Non-interface methods] Calling regular non interface methods using object literal ie non addressable _ with and without pointer receiver")

	fmt.Println("--> Will run method using object literal (ie non addressable) on a pointer receiver method")

	user{"JD"}.testNoPointerReceiver()
	// !!works since directly using user object here without creating a variable is not addressable, so temporary value. A value is passed as a copy to function
	// which has no pointer receiver but a value receiver, so value on value will work

	// user{"JD"}.testWithPointerReceiver()
	// !! won't work: cannot call pointer method on user literal
	// testWithPointerReceiver function has pointer receiver
	// user "literal" is a temporary value ie non addressable so compiler cannot implicitly set as (&user) so cannot work with pointer receiver

	fmt.Println()

	// (3) Calling regular non interface methods using variable set as pointer address _ calling function that has value receiver to showcase compiler dereferencing
	fmt.Println("(3) [Non-interface methods] Calling regular non interface methods using variable set as pointer address _ calling function that has value receiver to showcase compiler dereferencing")

	t := &user{"JD"}
	// set value as pointer address

	fmt.Printf("Type of user variable set as pointer address: %T\n", t)
	// type is *main.user

	t.testNoPointerReceiver()
	// will work even if non pointer receiver, compiler will implicitely derefernce it ie will load the value in receiver by doing (*t).testNoPointerReceiver()

	fmt.Println()

	// (4) Interface method: takes values that implement interface and send notifications
	fmt.Println("(4) [Interface method] Here you have to refer to table that shows method sets from the perspective of receiver types.")
	fmt.Println("Method here takes values that implement interface and send notifications")

	// sendNotif(u)
	// !! Error: This won't work
	// because from an interface perepective, if method receiver is pointer, values sent should be pointer type, which is not the case here we are sending value type not pointer type.

	sendNotif(&u) //This will work because from an interface perepective, if method receiver is pointer, values sent should be pointer type
}

// sendNotif accepts values that implement the notifier interface
// as you see from : n notifier
func sendNotif(n notifier) {
	n.notify()
}
