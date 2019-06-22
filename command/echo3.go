// Echo3.go a far simpler one line version of the prior programs
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//Use the join function to append all the arguments together
	fmt.Println(strings.Join(os.Args[1:], " "))

}
