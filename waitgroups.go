package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("Worker %d has started\n", i)
	time.Sleep(time.Duration(i) * time.Second) //This sleep simulates a heavy worker task
	fmt.Printf("Worker %d has done executing.\n", i)
}

/*
When we use go routines with closure the variable has to be passed to the anonymouse function
either by argument or by reinitialzing it to be accessible otherwise the func will have access
to the latest updated value of "i" as the loop breaks. This happens because for each iteration
goroutine refer to the same variable "i" and at the end of the loop it's value has been updated
to the latest one, that's why you see the weird output.
*/
func main() {
	var wg sync.WaitGroup

	// Lets say we have 5 worker tasks

	for i := 1; i < 6; i++ {
		wg.Add(1)
		// i := i0
		time.Sleep(time.Second)
		go func(i int) {
			fmt.Println("Inside: ", i)
			worker(i)
			defer wg.Done()
		}(i)

		// defer wg.Done()

	}
	wg.Wait()
}
