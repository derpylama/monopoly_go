package dice

import (
	"math/rand"
)

type Dice struct {
	numberOfDice int
	diceSides    int
}

func (d *Dice) throwDice() {
	var thrownDice []int

	for i := 0; i <= d.diceSides; i++ {
		thrownDice[i] = rand.Intn(6) + 1
	}

}
