package main

import (
	"fmt"
)

func main() {

	// 3 Ways to define a variable

	//Var Declared not defined
	var var1 int

	//Var declared as well as defined

	var var2 int = 1

	//Declare as well as defined dynamically (Short-Hand)

	var3 := "Test"

	//  Primitives
	/*
		int, uint8, uint32, uint64
		float, float32, float64
		bool (boolean)
		String
		rune



	*/

	fmt.Println(var1, var2, var3)
	// To show data type
	fmt.Printf("Type of var 1 is %T, var2 is %T, var3 is %T", var1, var2, var3)
}
