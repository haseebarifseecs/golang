package main

import (
	"fmt"
)

func main() {
	var array [5]int

	fmt.Print(array)

	array[2] = 1

	for i, j := range array {
		fmt.Println(i, j)
	}

	// Slices
	// Composite Literals x := []type {values}

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(slice[1:], slice[:len(slice)], slice[1:2])

	// Slice are like lists in python which allows dynamic array functions
	//Append

	slice = append(slice, 1, 2, 3, 4, 5)

	// Just like spread operator which allow us to append variable length slice to another slice

	slice2 := []int{1, 3, 4}

	slice = append(slice, slice2...) // ... at start means variable arguments numbers ... at the end means append variable length slice/values

	//You can also provide the max array capacity to slice using array.
	// make([]type, size, capacity)
	slice3 := make([]int, 10, 100)
	fmt.Println(len(slice3), cap(slice3), slice3)

	sliceone := []int{1, 2}
	slicetwo := []int{3, 4}
	sliceonetwo := [][]int{sliceone, slicetwo}

	fmt.Println(sliceonetwo)

}
