package game

import (
	"monopoly/events"
)

func (game *Game) RegisterListeners(commandChan chan<- GameCommand) {

	game.bus.Subscribe(events.StartTurn, func(e events.GameEvent) {
		p := e.Payload.(events.StartTurnPayload)
		game.bus.Publish(events.GameEvent{
			Type: events.InputPromptOptions,
			Payload: events.InputPromptPayload{
				PlayerName: p.PlayerName,
				Options:    []any{GameCommand{Type: CmdEndTurn, PlayerName: p.PlayerName}, GameCommand{Type: CmdRollDice, PlayerName: p.PlayerName}},
			},
		})
	})

	game.bus.Subscribe(events.PaidTax, func(ge events.GameEvent) {
		p := ge.Payload.(events.TaxPayload)

		game.bus.Publish(events.GameEvent{
			Type: events.UpdateMoney,
			Payload: events.UpdateMoneyPayload{
				PlayerName: p.PlayerName,
				Money:      p.PlayerMoney,
			},
		})
	})
}

func (game *Game) RegisterPromptListeners(commandChan chan<- GameCommand) {
	game.bus.Subscribe(events.InputPromptRollDice, func(e events.GameEvent) {
		p := e.Payload.(events.PromptRollDicePayload)

		commandChan <- GameCommand{
			Type:       CmdRollDice,
			PlayerName: p.PlayerName,
		}

	})
}
