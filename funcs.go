package main

import (
	"fmt"
)

// func name(arg type) returntype {}
func foo(s string) string {
	return s
}

// Variadic Parameter
/*
Variadic Parameter Converts Arguments into a slice of
length == no. of arguments passed
*/
func anything(s ...string) []string {
	fmt.Printf("%T\t", s)
	return s
}

/*

... variadic function must follow the variadic paramter as the last parameter to make that argument kind of optional but
if the parameter isn't the last element then it will complain as it must be the last parameter.
func m1 (s string, i ...int) {} will work if i dont pass any int
func m2 (s string..., i int) {} will not work


*/

func first() {
	fmt.Println("First Called")
}

func second() {
	fmt.Println("Second Called")
}

func third() {
	fmt.Println("Third Called")
}
func m1(s string, i ...int) string {
	return s
}

func main() {
	fmt.Println(foo("hello"))
	fmt.Println(anything("1", "2", "3"))
	// Let's say we have slice of strings

	s := []string{"1", "2", "3"}
	// If I directly pass the slice to the anything function it will raise exception as it's expecting a string but we passed slice of string
	// So we need to use variadic operator in reverse to unfurl it think of it like exporting values out of slice and passing them all

	fmt.Println(anything(s...)) // fmt.Println(anything(s)) --> This will not work
	/*
		defer keyword allow us to make sure the function is always executed at the end
		it uses stack to store the call, and always executes the one called at the last


	*/
	defer first()
	defer second()
	defer third() // this should run before first

	// Anonymouse Function

	func(s string) {
		fmt.Println(s + "called")
	}("Anonymouse")
}
