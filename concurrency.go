// Concurrency and parallelism are not the same
// Concurrency means structuring program so that it may run in parallel. (Speed Up Execution)
// Unlike parallelism Concurrent Program can execute on single thread
// There may be thousands of concurrent programs running on single thread.
// We use goroutines to achieve concurrency
// Goroutines are lightweight threads

package main

import (
	"fmt"
	"time"
)

func numSearch(targetVal int) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	time.Sleep((time.Duration(targetVal) * 100) * time.Millisecond)
	for _, i := range nums {
		if i == targetVal {
			fmt.Println("Found")
			return
		}
	}
	fmt.Println("Not Found.")
}
func main() {
	// fmt.Println("First Call Searching for 10")
	// numSearch(10) // This will wait 10 * 100 = 1000 ms

	// fmt.Println("Second Call Searching for 1")
	// numSearch(1) // This will only wait for 1 * 100 = 100ms
	//You may have notice 10 takes longer as it has to sleep for 1000ms while 1 executes faster but still 10 is executed before 1 due to program executing in sequential

	// Now if we use go routines

	go numSearch(10)
	go numSearch(1)

	/*
		The reason it doesn't printed the output of
		func calls and just printed the last statement and
		exited is because when you use goroutine it doesn't wait
		for the code to finish before moving on the next line
		it just calls and continue the execution flow. Since main
		function exited earlier then the function output that's why it
		doesn't printed the output as the compiler completed the execution
		after main
	*/

	// If i add delay longer then functions it will print out the output

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Completed")
}
