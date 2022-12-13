package main

import (
	"fmt"
)

/*
func (receiver type) name (args type) returntype { body }

# For Struct Methods

Receiver of type Non Pointer can take Pointer as well as Non Pointer
Receiver of type Pointer can only take Pointer
*/
type shape interface {
	area() int
	areaP() int
}
type rectangle struct {
	width  int
	length int
}

func (r rectangle) area() int {
	return (r.width * r.length)
}

func (r *rectangle) areaP() int {
	return r.width * r.length
}

func calcArea(r rectangle) int {
	return r.areaP()
}

// Always use var *a = &a can't do var a = &a
// * Dereference & -> Points to memory Address
func main() {
	a := 10
	b := &a
	fmt.Printf("%T \t \t \n %v", b, b)
	fmt.Println("\nABC", *&a)

	// This will work
	var r3 rectangle = rectangle{
		width:  1,
		length: 1,
	}
	r1 := rectangle{10, 10}
	fmt.Println(calcArea(r1))

	// r2 := &rectangle{5, 5}
	// fmt.Println(calcArea(r2))

	// fmt.Println(calcArea(r2))

	fmt.Println(calcArea(r1))
	// This will not work

	fmt.Println(calcArea(r1))

	fmt.Println(calcArea(r3))
}
