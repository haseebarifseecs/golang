package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i

		}
		close(c)
	}()
	// close(c) Above is async it will close channel first before it gets the chance to add the value

	for i := range c { //It will start pulling off the values from the channel
		// Range will not pull values till the channel is close if channel is not closed it will throw error
		fmt.Println(i)
	}

}
