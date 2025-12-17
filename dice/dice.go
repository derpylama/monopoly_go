package dice

import (
	"math/rand"
)

type Dice struct {
	numberOfDice int
	diceSides    int
}

func (d *Dice) ThrowDice() []int {
	thrownDice := make([]int, d.numberOfDice)

	for i := 0; i < d.numberOfDice; i++ {
		thrownDice[i] = rand.Intn(6) + 1
	}

	return thrownDice
}

func NewDice(numberOfDice int, diceSides int) *Dice {
	return &Dice{
		numberOfDice: numberOfDice,
		diceSides:    diceSides,
	}
}
