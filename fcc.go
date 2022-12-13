package main

import (
	"fmt"
)

// In Go funcs are first class citizens

func f2() func() {
	return func() {
		fmt.Println("No Return for you")
	}
}

func f3() func() int {
	return func() int {
		return 1
	}
}

func f4(f func() func()) {
	f()()
}

func oddSum(num ...int) int {
	var sum int
	for _, j := range num {
		sum += j
	}

	return sum
}

func odd(f func(num ...int) int, num ...int) int {
	var o []int
	for _, j := range num {
		if j%2 != 0 {
			o = append(o, j)
		} else {
			continue
		}
	}
	sum := f(o...)

	return sum
}
func main() {
	f1 := func() {
		fmt.Println("f1 Called")
	}
	f2()()
	f1()
	f2 := f2()
	f2()

	f3 := f3()
	fmt.Println(f3())
	f4(
		func() func() {
			return func() {
				fmt.Println("callback")
			}
		},
	)
	fmt.Println(odd(oddSum, 1, 2, 3))

}
