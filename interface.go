package main

import (
	"fmt"
)

// Interfaces allow to make methods of different types (Polymorphism)

/*
Since we have an interface of type Uni which has a method teach
both structs teacher and department implements the method teach()
we can say both teacher and department are of type uni since they
teach in departments and departments are part of it
*/
type teacher struct {
	name       string
	department department
}

type department struct {
	name string
}

type uni interface {
	teach()
}

func (t teacher) teach() {
	fmt.Println(t.name + "is teaching")
}

func (d department) teach() {
	fmt.Println(d.name + "is teaching.")
}

func print(u uni) {
	fmt.Println(u)

}

func determineType(u uni) {
	switch u.(type) {
	case teacher:
		fmt.Println("Type is teacher", u.(teacher).name)
	case department:
		fmt.Println("Type is department", u.(department).name)
	default:
		fmt.Println("Type is unknown")
	}

}
func main() {
	t := teacher{
		name: "XYZ",
		department: department{
			name: "department",
		},
	}

	d := department{
		name: "department",
	}
	// Since both structs are indirectly part of interface uni i can pass their objects as they implemented the teach method. (Override)
	print(t)
	print(d)
	determineType(t)

}
