package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan bool) {
	fmt.Printf("Worker %d has started.\n", i)
	time.Sleep(time.Duration(i) * time.Millisecond)
	fmt.Printf("Worker %d has stopped.\n", i)
	c <- true
}
func main() {
	c := make(chan bool)
	for i := 1; i < 6; i++ {
		go worker(i, c)
	}
	// Without Receiver it will print nothing
	for i := 1; i < 6; i++ {
		<-c
	}
}
