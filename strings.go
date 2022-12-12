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
}
