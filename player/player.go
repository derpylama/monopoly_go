package player

import (
	"fmt"
	cards "monopoly/card"
	"monopoly/helper"
)

type Player struct {
	position        int
	money           int
	name            string
	cardsInHand     []cards.Card
	isJailed        bool
	jailedTurns     int
	amountOfDubbles int
	hasRolled       bool
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

func (player *Player) ToggleIsJailed() {
	if player.isJailed {
		player.isJailed = false
	} else {
		player.isJailed = true
	}
}

func (player Player) HasRolled() bool {
	return player.hasRolled
}

func (player Player) GetJailedTurns() int {
	return player.jailedTurns
}

// Incremeants the players turns spent in jail to a max of 3
func (player *Player) IncrementJailedTurns() {
	jailTurns := player.jailedTurns
	jailTurns++
	player.jailedTurns = helper.Clamp(jailTurns, 0, 3)
}

func (player *Player) ResetJailedTurns() {
	player.jailedTurns = 0
}

func (player *Player) GetHasRolled() bool {
	return player.hasRolled
}

func (player *Player) ToggleHasRolled() {
	if player.hasRolled {
		player.hasRolled = false
	} else {
		player.hasRolled = true
	}
}

func (player *Player) IncrementAmountOfDubbles() {
	dubbles := player.amountOfDubbles
	dubbles++
	player.amountOfDubbles = helper.Clamp(dubbles, 0, 3)
}

func (player *Player) GetAmountOfDubbles() int {
	return player.amountOfDubbles
}

func (player *Player) SetMoney(money int) {
	player.money = helper.Clamp(money, 0, 100000000)
}

func (player *Player) CanAfford(amount int) bool {
	return player.GetMoney() >= amount
}

func (player *Player) Move(rolledDice []int) {
	if !player.GetJailStatus() {
		curPos := player.GetPosition()

		var diceTotal int

		for _, dice := range rolledDice {
			diceTotal += dice
		}

		newPos := (diceTotal + curPos) % 39

		if newPos < curPos {
			//Passed GO
			player.SetMoney(player.GetMoney() + 200)
		}

		player.position = newPos
	}
}

func (player *Player) PayRent(playerToPay *Player, amount int) {
	if player.Pay(amount) {
		playerToPay.SetMoney(playerToPay.GetMoney() + amount)
	} else {
		fmt.Println("You can't afford to pay the rent")
	}
}

func (player *Player) Pay(cost int) bool {
	if player.GetMoney() >= cost {
		player.SetMoney(player.GetMoney() - cost)
		return true
	}

	return false
}

func (player *Player) Teleport(pos int) {

	newPos := helper.Clamp(pos, 0, 39)

	player.position = newPos
}

func NewPlayer(money int, name string) *Player {
	return &Player{
		position:        0,
		money:           money,
		name:            name,
		cardsInHand:     []cards.Card{},
		isJailed:        false,
		hasRolled:       false,
		amountOfDubbles: 0,
	}
}
