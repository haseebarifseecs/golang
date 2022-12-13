/*
	In Go everything is a type. Likewise, Error is also a type.
	Error is of type interface which has a method Error() string.
	So, if every struct implements the method Error() string, is also of
	type error. (Implicitlly)

*/

package main

import (
	"fmt"
)

func main() {
	n, er := fmt.Println("") //FMT Println package returns nummber of lines and a error descriptor
	if er != nil {
		fmt.Print(er)
	}

	fmt.Print(n)

	var (
		name string
		age  string
		role string
	)

	fmt.Println("Name:")
	{
		_, err := fmt.Scan(&name)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Age:")
	{
		_, err := fmt.Scan(&age)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("role")
	{
		_, err := fmt.Scan(&role)

		if err != nil {
			panic(err)
		}
	}

}
