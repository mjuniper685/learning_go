// A program building a simple ticketing system for Mars bound rockets

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//var spaceAdv, spaceX, virgin = "Space Adventures", "Space X", "Virgin Galactic"
	// rand.Intn produces a random number between 0 and n - 1
	// seeding the number means we don;t get the same value each time!
	rand.Seed(time.Now().UnixNano())
	var spacelinerand = rand.Intn(3) + 1
	var spaceline string
	//var distance int = 62100000
	var minSpeed = 16
	var maxSpeed = 30

	var count = 10

	//define table
	fmt.Printf("%-15v %6v %6v %6v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=========================================")

	for count > 0 {
		fmt.Println(spacelinerand)
		// pick a spaceline
		switch spacelinerand {
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "Space X"
		case 3:
			spaceline = "Virgin Galactic"
		}

		fmt.Println(spaceline)
		speed := rand.Intn(maxSpeed-minSpeed) + minSpeed
		fmt.Println("Speed = ", speed)
		spacelinerand = rand.Intn(3) + 1
		count--

	}

}
