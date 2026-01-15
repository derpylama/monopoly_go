package logger

import (
	"fmt"
	"monopoly/common"
	"monopoly/events"
	"monopoly/game"
	"strconv"
)

func RegisterListeners(bus *common.Bus, l *Logger) {

	bus.Subscribe(common.LandedOnFreeParking, func(e common.GameEvent) {
		p := e.Payload.(events.LandedOnFreeParkingPayload)
		l.log("Player " + p.PlayerName + " landed on Free Parking")
	})

	bus.Subscribe(common.RolledDice, func(e common.GameEvent) {

		p := e.Payload.(events.RolledDicePayload)
		fmt.Println("DEBUG: RolledDice received")
		l.log(
			p.PlayerName +
				" rolled " + strconv.Itoa(p.Dice[0]) + " and " + strconv.Itoa(p.Dice[1]))

	})

	bus.Subscribe(common.PaidRent, func(e common.GameEvent) {
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

	bus.Subscribe(common.PaidTax, func(e common.GameEvent) {
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

	bus.Subscribe(common.StartTurn, func(e common.GameEvent) {
		p := e.Payload.(events.StartTurnPayload)
		l.log("It's now " + p.PlayerName + "'s turn!")
	})
}

func RegisterPromptListeners(bus *common.Bus, commandChan chan<- game.GameCommand) {

}
