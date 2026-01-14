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
		fmt.Println("DEBUG: RolledDice received")
		l.log(
			p.PlayerName +
				" rolled " + strconv.Itoa(p.Dice[0]) + " and " + strconv.Itoa(p.Dice[1]))

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

	// bus.Subscribe(events.InputPromptOptions, func(e events.GameEvent) {
	// 	p := e.Payload.(events.InputPromptPayload)
	// 	l.log("Prompting " + p.PlayerName + " for all inputs")
	// 	for _, Options := range p.Options {
	// 		l.log("- " + Options)
	// 	}

	// })

	bus.Subscribe(events.StartTurn, func(e events.GameEvent) {
		p := e.Payload.(events.StartTurnPayload)
		l.log("It's now " + p.PlayerName + "'s turn!")
	})
}

func RegisterPromptListeners(bus *events.Bus, commandChan chan<- game.GameCommand) {

}
