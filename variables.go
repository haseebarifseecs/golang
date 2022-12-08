package main

import (
	"fmt"
)

// Public Global Variable must have First Letter Capitalized
const Public string = "This is publicly accessible constant"

func main() {

	// 3 Ways to define a variable

	//Var Declared not defined
	var var1 int

	//Var declared as well as defined

	var var2 int = 1

	//Declare as well as defined dynamically (Short-Hand)

	var3 := "Test"

	var var4 = 1

	//This assignment will fail var4 = 2 because go is strictly typed

	//  Primitives
	/*
		int, uint8, uint32, uint64
		float, float32, float64
		bool (boolean)
		String
		rune



	*/

	fmt.Println(var1, var2, var3, var4, Public)
	// To show data type

	var3 = "Something else"
	fmt.Print(var3 + "\n")
	fmt.Printf("Type of var 1 is %T, var2 is %T, var3 is %T \n", var1, var2, var3)
}
