// Go file explaining maps
// maps map keys to values

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// define a map as Vertex of strings
var m map[string]Vertex

var n = map[string]Vertex{
	"Microsoft": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	//
	m = make(map[string]Vertex)
	// set key/value Vertex in the map
	// This is a map literal
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	// print value at key Bell Labs
	fmt.Println(m["Bell Labs"])
	fmt.Println(n)
}
