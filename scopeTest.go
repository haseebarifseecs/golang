package main

import (
	"fmt"

	s "package.mod/scope"
)

func main() {
	fmt.Print(s.GlobalVar) //accessible
	// fmt.Print(s.packageLevel) //not accessible (different package)
	// fmt.Print(s.local)        //not accessible
}
