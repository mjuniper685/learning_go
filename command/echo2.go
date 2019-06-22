// Echo2 prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// the _ is a blank identifier used when syntax requires a var
	// but logic does not, so her range is used instead of a var
	// also using a slice [1:] instead of a len
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
