// A Go program exploring Slices

package main

import "fmt"

func main() {
	// create a dynamically sized array using make
	a := make([]int, 5) //len(a) == 5
	printSlice("a", a)

	// to specify a capacity, pass a 3rd argument to make
	b := make([]int, 0, 5) //len(b) == -, cap(b) == 5
	printSlice("b", b)

	// using the append function
	// func append(s []T, vs ...T) []T
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// adding more than one element
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
