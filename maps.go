package main

import (
	"fmt"
)

func main() {
	// Maps are like hashmaps in Go which store key value pairs.

	// m := map[keytype] valuetype { key1 : value1}

	m := map[string]string{
		"name": "N/A",
		"age":  "N/A",
	}

	fmt.Println(m["name"])

	// It will print false for key that doesn't exist
	// Value, isExist := m["name"] if key exists it will print out the value as well as the boolean if not it will return default value of the type you chose
	k, v := m["name"]
	fmt.Println(k, v)

	k1, v1 := m["notexist"]
	fmt.Println(k1, v1)

	if _, v1 := m["notexist"]; v1 == false {
		fmt.Println("Doesn't exist")
	}
	// Iterating

	for key, value := range m {
		fmt.Println(key, value)
	}

	// Deleting Key in map
	// delete(map, key)

	delete(m, "name")

}
