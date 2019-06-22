// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.package duplicates
// Use Ctrl+D when running this to send EOF to the input.Scan() and exit the loop

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// set of k/v pairs strings/int in this case
	// make creates a new empty map
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// get the input into the map as a key
		// increment the value
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
