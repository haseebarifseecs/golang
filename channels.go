// Channels allow us to receive values from multiple goroutines
// Comms b/w goroutines
// Both sending and receiving in channels are blocking, So Compiler will wait until it receive the output to finish
// Channels block if both send and receive doesn't happen at the same time.
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runner(s string, c chan string, slice chan []int, m chan map[string]string) string {
	time.Sleep(time.Duration(rand.Intn(10)))
	c <- s
	intSlice := []int{123}
	slice <- intSlice
	sampleMap := map[string]string{
		"1": "1",
	}
	m <- sampleMap
	return s
}
func main() {
	fmt.Println("")
	sliceChannel := make(chan []int)
	mapChannel := make(chan map[string]string)
	channel := make(chan string)
	helloworld := []string{"hello", "world"}

	for _, i := range helloworld {
		go runner(i, channel, sliceChannel, mapChannel)
	}

	for i := 0; i < 2; i++ {
		fmt.Println(<-channel)
		fmt.Println(<-sliceChannel)
		fmt.Println(<-mapChannel)

	}

	/*

		By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
		Here we make a channel of strings buffering up to 2 values.
		Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	*/
	bufferedChannel := make(chan string, 2)
	bufferedChannel <- "First"
	bufferedChannel <- "Second"

	fmt.Println(<-bufferedChannel) //First
	fmt.Println(<-bufferedChannel) //Second

}
