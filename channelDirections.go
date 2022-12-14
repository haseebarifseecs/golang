package main

import (
	"fmt"
)

func onlySends(c chan<- string, s string) {
	c <- s
}

func onlyReceives(c <-chan string) {
	fmt.Println(<-c)
}

//	func thisErrors(c chan<- string, s string) error {
//		// Cant receive to a sending channel
//		_, err := fmt.Println(<-c)
//		if err != nil {
//			return errors.New("Error!")
//		}
//		return nil
//	}
func main() {
	fmt.Println("You can specify channel specificity to whether send or receive")
	c := make(chan string, 1)
	defer close(c)
	onlySends(c, "Hello World!")
	onlyReceives(c)
}
