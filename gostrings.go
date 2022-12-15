package main

import (
	"fmt"
	"strconv"
	"strings"
)

func overlap(s string, sub string) int {
	count := 0
	length := len(sub)
	window := 0
	if length > len(s) {
		return 0
	}
	for i := length; i <= len(s); i++ {
		// fmt.Println("Printing substr \n", string(s[window:i]))
		if string(s[window:i]) == sub {
			count++
		}
		window++
	}
	return count

}
func main() {
	// Since Strings in go are byte array we can store them in a byte

	var s1 byte
	s1 = 'A'
	fmt.Println(s1) //It will print Byte Value

	fmt.Printf("%#U\n", s1) //Unicode + Actual Character

	var s2 string
	s2 = string(65) // It will convert ASCII 65 to its string value

	fmt.Println(s2) // Should Print "A"

	// Type Conversion

	var s3 string
	s3 = strconv.Itoa(65) // This will convert int 65 to string '65'
	fmt.Println(s3)
	// Converting String to Int

	var s4 int
	s4, _ = strconv.Atoi("65")
	fmt.Println(s4)

	fmt.Printf("Type of s4 is %T \n", s4)

	var s5 float64

	s5, _ = strconv.ParseFloat("2.1", 32)

	fmt.Printf("Type of S5 is %T\n", s5)

	s6 := "hello world"
	fmt.Println(strings.Count(s6, "o"))

	// Substrings

	fmt.Println(s6[3:len(s6)])

	firstChar := s6[0]
	if firstChar == 'h' {
		fmt.Println("h is matched")
	}

	// Check if a substr occur in string

	s7 := "hello"
	fmt.Println(strings.Contains(s7, "e"))

	s8 := "CGATATA"
	fmt.Println("Overlap Output: \t", overlap(s8, "ATA")) // should output 2

	s8 = "ATATATATATATA"
	fmt.Println("Overlap Output: \t", overlap(s8, "ATA")) // should output 6

	s8 = "SASASQW"
	fmt.Println("Overlap Output: \t", overlap(s8, "ATA")) // should output 0

	// Convert to Uppercase

	s8 = "lower"
	fmt.Println(strings.ToUpper(s8))

	// Lower
	fmt.Println(strings.ToLower(s8))

	// Has Prefix

	fmt.Println(strings.HasPrefix(s8, "lower"))

	// Has Suffix
	fmt.Println(strings.HasSuffix(s8, "lower"))

	fmt.Println(strings.Cut(s8, "ow"))

	fmt.Println(strings.Replace(s8, "lower", "upper", 1))

	var stringBuilder strings.Builder

	stringBuilder.Write([]byte("Hello"))
	fmt.Println(stringBuilder.Cap()) //Prints capacity 2^n 0,2,4,8...
	stringBuilder.Len()              //Length

	stringBuilder.Grow(5) //Grow the length 2*currentCap + 5
	val := stringBuilder.String()

	fmt.Println("String Builder Val", val)

	fmt.Println(stringBuilder.Cap())
}
