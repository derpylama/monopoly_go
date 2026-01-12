package logger

import (
	"fmt"
	"monopoly/events"
	"monopoly/game"
	"strconv"
)

func RegisterListeners(bus *events.Bus, l *Logger) {

	bus.Subscribe(events.LandedOnFreeParking, func(e events.GameEvent) {
		p := e.Payload.(events.LandedOnFreeParkingPayload)
		l.log("Player " + p.PlayerName + " landed on Free Parking")
	})

	bus.Subscribe(events.RolledDice, func(e events.GameEvent) {
		p := e.Payload.(events.RolledDicePayload)
		l.log(
			p.PlayerName +
				" rolled " +
				fmt.Sprintf("%d + %d", p.Dice[0], p.Dice[1]),
		)
	})

	bus.Subscribe(events.PaidRent, func(e events.GameEvent) {
		p := e.Payload.(events.PaidRentPayload)
		l.log(
			p.PlayerName +
				" paid $" +
				strconv.Itoa(p.Rent) +
				" rent to " +
				p.Owner +
				" for " +
				p.TileName,
		)
	})

	bus.Subscribe(events.PaidTax, func(e events.GameEvent) {
		p := e.Payload.(events.TaxPayload)
		l.log(
			p.PlayerName +
				" paid $" +
				strconv.Itoa(p.TaxAmount) +
				" in tax at " +
				p.TileName,
		)
	})
}

func RegisterPromptListeners(
	bus *events.Bus,
	commandChan chan<- game.GameCommand,
) {

	bus.Subscribe(events.InputPromptRollDice, func(e events.GameEvent) {
		p := e.Payload.(events.RolledDicePayload)

		fmt.Printf("%s, type 'roll' to roll the dice: ", p.PlayerName)

		var input string
		fmt.Scanln(&input)

		if input == "roll" {
			commandChan <- game.GameCommand{
				Type:       game.CmdRollDice,
				PlayerName: p.PlayerName,
			}
		}
	})
}
