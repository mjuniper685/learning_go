// Code file explaining ranges

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range for loop iterates over a slice or map
	// For a slice two values are returned the index and the element at that index
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// same again but assigning index to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
