package main

import (
	"fmt"
)

// iota allow to automatically increment values by one

const (
	a = iota
	b
	c
	d
	e
	f
	g
)

const (
	_  = iota             // iota = 0 (ignored)
	kb = 1 << (iota * 10) // iota = 1
	mb = 1 << (iota * 10) // iota = 2
	gb = 1 << (iota * 10) // iota = 3
	tb = 1 << (iota * 10) // iota = 4
)

func main() {
	// Strings are immutable and are represented as sequence of bytes
	// byte is same as uint8 in Go
	fmt.Println("\n", kb/1024, mb/1024, gb/1024, tb/1024)
	var (
		q   = "one"
		w   = "two"
		e_e = "three"
	)

	fmt.Print(q, w, e_e)
	fmt.Println(a, b, c, d, e, f, g)
	var a string = "Bytes"
	var byteString []byte = []byte(a)

	fmt.Println(a)

	fmt.Println(byteString)

	fmt.Printf("%x ", a)

	for i := 0; i < len(byteString); i++ {
		fmt.Printf("%#x", byteString[i])
	}

	// same is for i,j in enumerate(iterator)
	for i, j := range byteString {
		fmt.Print(i, j)
	}

	// while i < 10
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}

	var k float64 = 43 / 40
	kkk := float64(43) / 40
	kk := 43 / 40

	fmt.Println("\n", k, kk, kkk)

	i = 0

	for i < 100 {
		if i%2 == 0 {
			fmt.Println("\nEven", i)
		} else {
			break
		}
		i++

	}

	i = 0

	for {
		if i%2 == 0 {
			fmt.Println(i)
		}
		if i == 10 {
			break
		}
		i++
	}

	// Printing ascii 33 - 122
	for p := 0; p <= 122; p++ {
		fmt.Printf("%#U\n", p)
	}

	isTrue := true

	if isTrue {
		fmt.Print(isTrue)
	}

	// isTrue still accessible here if i do fmt.Print(isTrue) here but if I want to restrict variable scope then i should do this

	if isFalse := false; isFalse {
		fmt.Print("Not Running")
	}

	// fmt.Print(isFalse) Not Accessible
	{
		var (
			abc = "abc\n\n\n"
		)
		fmt.Print(abc) //Only Accessible in this block
	}
	// fmt.Println(abc) Not Accessible

	var switchVar = true
	switch switchVar {
	case (switchVar == false):
		fmt.Println("Not Printing")
	case (switchVar == true), (false): //In Golang you can only compare similar data types for instance you can't compare bool to string or string to bool otherwise it will throw error
		fmt.Println("Printing")
		fallthrough
	case (true):
		fmt.Println("Will not print without fallthrough")
		fallthrough //Since I add fallthrough here it will print the next statement whether it's true or false
	case (false):
		fmt.Print("Will not Print")

	default:
		fmt.Print("Will only print if nothing is true")
	}
}
