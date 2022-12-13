package main

import (
	"errors"
	"fmt"
	"log"
)

func throwFatal() {
	log.Fatal("Something went fatal")
}
func throwException() {
	panic("Oh No! There is something wrong")
}

type e struct {
	val string
}

func (e *e) something() (string, error) {
	if e.val == "" {
		return "", errors.New("Something went wrong")
	} else {
		return "hello", nil
	}

}

func main() {
	// Panic Stops the execution of program at the point where it occurs
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Panic", r)
		}
	}()
	throwFatal() //This will even prevent defer functions from running
	throwException()

	fmt.Println("Completed!")
	e := e{""}
	a, err := e.something()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	// We also have fmt.Errorf("") option which allow us to format the error which is not available in errors.New() which just accept the string

	fmt.Errorf("Some error %#U", "ABC")

}
