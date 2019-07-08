// Structs in Go

package main

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

type Employee struct {
	firstName string
	lastName  string
	salary    int
	age       int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// Access the fields using a dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// You can also access them through pointers
	p := &v
	p.X = 1e9
	fmt.Println(v)

	emp1 := Employee{
		firstName: "Sam",
		lastName:  "Sparrow",
		age:       25,
		salary:    25000,
	}

	fmt.Println(emp1)

}
