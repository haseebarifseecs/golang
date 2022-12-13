package main

import (
	"fmt"
)

type rectangle struct {
	width  int
	length int
}

type shape interface {
	areaP() int
}

func (r *rectangle) areaP() int {
	return r.width * r.length
}

func calcArea(r rectangle) int {
	return r.areaP()
}

func main() {
	r1 := rectangle{10, 10}
	fmt.Println(calcArea(r1))

}
