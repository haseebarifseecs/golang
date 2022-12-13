package main

import (
	"fmt"
)

// Anon structs are used to keep the scope clean
func main() {
	anon := struct {
		value string
	}{
		value: "This struct is anonymouse it doesn't have any name",
	}

	fmt.Println(anon)

	anon.value = "Changed"

	fmt.Println(anon)
}
