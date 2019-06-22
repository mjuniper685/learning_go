// Echo1.go command line formatting
package main

import (
	"fmt"
	"os"
)

func main() {
	// build up a string using a for loop and appending
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
