package player

import (
	cards "monopoly/card"
	"monopoly/helper"
)

type Player struct {
	position    int
	money       int
	name        string
	cardsInHand []cards.Card
	isJailed    bool
}

func (player Player) GetPosition() int {
	return player.position
}

func (player Player) GetMoney() int {
	return player.money
}

func (player Player) GetName() string {
	return player.name
}

func (player Player) GetCardsInHand() []cards.Card {
	return player.cardsInHand
}

func (player Player) GetJailStatus() bool {
	return player.isJailed
}

func (player *Player) SetMoney(money int) {
	player.money = helper.Clamp(money, 0, 100000000)
}

func (player *Player) Move(rolledDice []int) {
	curPos := player.GetPosition()

	var diceTotal int

	for _, dice := range rolledDice {
		diceTotal += dice
	}

	newPos := (diceTotal + curPos) % 40

	player.position = newPos
}

func (player *Player) Pay(cost int) bool {
	if player.GetMoney() >= cost {
		player.SetMoney(player.GetMoney() - cost)
		return true
	}

	return false
}

func NewPlayer(money int, name string) *Player {
	return &Player{
		position:    0,
		money:       money,
		name:        name,
		cardsInHand: []cards.Card{},
		isJailed:    false,
	}
}
