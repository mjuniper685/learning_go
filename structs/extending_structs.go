// Extending structs using interfaces

package main

import (
	"fmt"
)

type attacker struct {
	attackpower int
	dmgbonus    int
}

// contains sub-struct attacker
type sword struct {
	attacker  // don't need to type attacker again as name same as type
	twohanded bool
}

type gun struct {
	attacker
	bulletsremaining int
}

func main() {

	sword1 := sword{attacker: attacker{attackpower: 1, dmgbonus: 5}, twohanded: true}
	gun1 := gun{attacker: attacker{attackpower: 10, dmgbonus: 20}, bulletsremaining: 11}

	fmt.Printf("Weapons: sword: %v, gun: %v\n", sword1, gun1)

	sword1.Wield()

	gun1.Wield()

}

// using the method syntax to extended gun and sword with methods
func (s sword) Wield() bool {
	fmt.Println("You've wielded a sword!")
	return true
}

func (g gun) Wield() bool {
	fmt.Println("You've wielded a gun!")
	return true
}
