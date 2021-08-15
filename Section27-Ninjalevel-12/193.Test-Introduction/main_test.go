package main

import "testing"

//File name needn't match as well.
//It can be mains_test.go as well. but main_test.go is the best practise.
//func TestAnyName(t *testing.T) // This is valid, fucntion name in test
//needn't match the function name we are testing.
func TestMySum(t *testing.T) { // this is the best practise.
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
