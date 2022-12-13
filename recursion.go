package main

import "fmt"

/*

	Recursion is a function calling itself again and again till it reaches a base condition.
	Each function call and its return is store in a Stack (FILO or LIFA)
	Once the Base Condition is met the last call() is return first and then escalates

*/

func factorial(x int) int {
	var fact int
	if x == 1 || x == 0 {
		return 1
	}

	fact = x * factorial(x-1)

	return fact

}

func main() {
	fmt.Println(factorial(3)) //6
}
