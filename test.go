package main

import (
	"fmt"

	e "package.mod/exercise"
)

var A string = "Abc"
var B string

func main() {
	a := "abc"
	remainder := e.RemainingOvenTime(30)
	var i interface{} = e.PreparationTime(2)
	elapsedtime := e.ElapsedTime(3, 20)

	fmt.Print(remainder, i, elapsedtime)
}
