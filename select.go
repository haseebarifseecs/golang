// Consider select as await and go + channels as async so Select will wait till all the channels receivers receive the values
// Select is useful like when you want to return as soon as any one of the multiple goroutine as completed
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "Data for C1"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		c2 <- "Data for C2"

	}()

	select {
	case d1 := <-c1:
		fmt.Println("Received:", d1)
	case d2 := <-c2:
		fmt.Println("Received: ", d2)
	}
	fmt.Println("Executed Completely")
}
