package main

import (
	"fmt"
)

// Concept of closure refers to keeping the context of variable scope as close as possible
// In below function scope of x is only available to the function increment and the nested function but with eachcall
// Variable is not redeclared instead its value keep on incrementing
func increment() func() {
	var x int
	return func() {
		x++
		fmt.Println(x)
	}
}

func main() {
	f1 := increment()
	f2 := increment()

	f1() // 1
	f2() // 1
	f1() // 2
	f1() // 2
}
