package game

import (
	"monopoly/common"
	"monopoly/events"
)

func (game *Game) RegisterListeners(commandChan chan<- GameCommand) {

	game.bus.Subscribe(common.StartTurn, func(e common.GameEvent) {
		p := e.Payload.(events.StartTurnPayload)
		game.bus.Publish(common.GameEvent{
			Type: common.InputPromptOptions,
			Payload: events.InputPromptPayload{
				PlayerName: p.PlayerName,
				Options:    []any{GameCommand{Type: CmdEndTurn, PlayerName: p.PlayerName}, GameCommand{Type: CmdRollDice, PlayerName: p.PlayerName}},
			},
		})
	})

	game.bus.Subscribe(common.PaidTax, func(ge common.GameEvent) {
		p := ge.Payload.(events.TaxPayload)

		game.bus.Publish(common.GameEvent{
			Type: common.UpdateMoney,
			Payload: events.UpdateMoneyPayload{
				PlayerName: p.PlayerName,
				Money:      p.PlayerMoney,
			},
		})
	})

	game.bus.Subscribe(common.Jailed, func(ge common.GameEvent) {
		player := game.getPlayer()
		player.IncrementJailedTurns()
	})
}

func (game *Game) RegisterPromptListeners(commandChan chan<- GameCommand) {
	game.bus.Subscribe(common.InputPromptRollDice, func(e common.GameEvent) {
		p := e.Payload.(events.PromptRollDicePayload)

		commandChan <- GameCommand{
			Type:       CmdRollDice,
			PlayerName: p.PlayerName,
		}

	})
}
