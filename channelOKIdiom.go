package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		c <- 1
	}()

	value, ok := <-c

	fmt.Println(value, ok)
	close(c)
	value, ok = <-c
	fmt.Println(value, ok)
}
