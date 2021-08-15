package main()
// Read https://blog.golang.org/defer-panic-and-recover
// Panic
// Panic is a built-in function that stops the ordinary flow of control and begins panicking. 
// When the function F calls panic, execution of F stops, 
// any deferred functions in F are executed normally, and then F returns to its caller. 
// To the caller, F then behaves like a call to panic. The process continues up the stack 
// until all functions in the current goroutine have returned, at which point the program crashes. 
// Panics can be initiated by invoking panic directly. 
// They can also be caused by runtime errors, such as out-of-bounds array accesses.


//Recover
// Recover is a built-in function that regains control of a panicking goroutine. 
// Recover is only useful inside deferred functions. <- VIMP
// During normal execution, a call to recover will return nil and have no other effect. 
// If the current goroutine is panicking, 
// a call to recover will capture the value given to panic and resume normal execution.

func main(){
	var x int
	x++
	fmt.Println(x)
	y := c()
	fmt.Prinln(y)

}

// I didn't understand this part.
// Deferred functions may read and assign to the returning function's named return values.
// In this example, a deferred function increments the return value i after the surrounding function returns.
// Thus, this function returns 2:
func c() (i int){
	defer func() {
		i++
	}()
	return 1
}

// // Defer funciton calls are executed in the Last In First Order after the surrounding fucntion returns.
// // This function returns "3210"
// func foo(){
// 	for i:=0;i<4;i++{
// 		defer fmt.Println(i)

// 	}
// }