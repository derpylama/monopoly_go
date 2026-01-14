package view

import (
	"fmt"
	"monopoly/events"
	"monopoly/game"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// func RegisterListeners(commandChan chan<- game.GameCommand, bus *events.Bus) {

// 	bus.Subscribe(events.StartTurn, func(e events.GameEvent) {
// 		p := e.Payload.(events.StartTurnPayload)

// 		bus.Publish(events.GameEvent{
// 			Type: events.InputPromptRollDice,
// 			Payload: events.PromptRollDicePayload{
// 				PlayerName: p.PlayerName,
// 			},
// 		})
// 	})

// }

func RegisterPromptListeners(commandChan chan<- game.GameCommand, bus *events.Bus) {
	// bus.Subscribe(events.InputPromptRollDice, func(e events.GameEvent) {
	// 	p := e.Payload.(events.PromptRollDicePayload)
	// 	fmt.Printf("%s, type 'roll' to roll the dice: ", p.PlayerName)
	// 	var input string

	// 	if input == "roll" {
	// 		commandChan <- game.GameCommand{
	// 			Type:       game.CmdRollDice,
	// 			PlayerName: p.PlayerName,
	// 		}
	// 	}
	// })
}

func RegisterReactiveGUIListeners(
	bus *events.Bus,
	g *game.Game,
	commandChan chan<- game.GameCommand,
	logArea *widget.Entry,
	buttonContainer *fyne.Container,
	playerLabel *widget.Label,
	mainWindow fyne.Window,
) {

	//Update which player's turn it is and their money
	bus.Subscribe(events.StartTurn, func(ge events.GameEvent) {
		payload := ge.Payload.(events.StartTurnPayload)

		fyne.Do(func() {
			playerLabel.SetText(
				fmt.Sprintf("%s's turn:\n%s has $%d", payload.PlayerName, payload.PlayerName, payload.Money),
			)
		})
	})

	//Update the players money after every time they spend it
	bus.Subscribe(events.UpdateMoney, func(ge events.GameEvent) {
		payload := ge.Payload.(events.UpdateMoneyPayload)

		fyne.Do(func() {
			playerLabel.SetText(
				fmt.Sprintf("%s's turn:\n%s has $%d", payload.PlayerName, payload.PlayerName, payload.Money),
			)
		})
	})

	// Show prompt options from the game
	bus.Subscribe(events.InputPromptOptions, func(e events.GameEvent) {
		payload := e.Payload.(events.InputPromptPayload)

		// Clear existing buttons
		buttonContainer.Objects = nil

		// Create new buttons for each option
		for _, opt := range payload.Options {
			optionName := opt.(game.GameCommand) // capture loop variable

			btn := widget.NewButton(string(optionName.Type), func() {
				fmt.Printf("Player %s chose %s\n", payload.PlayerName, optionName)
				commandChan <- game.GameCommand{
					Type:       game.CommandType(optionName.Type),
					PlayerName: payload.PlayerName,
				}
			})

			buttonContainer.Add(btn)
		}

		fyne.Do(func() { buttonContainer.Refresh() })

	})

	// Logging events
	bus.Subscribe(events.RolledDice, func(e events.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.RolledDicePayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s rolled %d + %d = %d\n",
				p.PlayerName, p.Dice[0], p.Dice[1], p.Dice[0]+p.Dice[1],
			))
		})
	})

	bus.Subscribe(events.LandedOnTile, func(e events.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.LandedOnTilePayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s landed on %s\n",
				p.PlayerName, p.TileName,
			))
		})
	})

	bus.Subscribe(events.PaidRent, func(e events.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.PaidRentPayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s paid $%d rent to %s for %s\n",
				p.PlayerName, p.Rent, p.Owner, p.TileName,
			))

		})
	})

	bus.Subscribe(events.PaidTax, func(e events.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.TaxPayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s paid $%d tax at %s\n",
				p.PlayerName, p.TaxAmount, p.TileName,
			))

		})
	})

	bus.Subscribe(events.LandedOnUnownedProperty, func(ge events.GameEvent) {
		fyne.Do(func() {
			p := ge.Payload.(events.LandedOnUnownedPropertyPayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s landed on unowned property %s (cost: $%d)\n",
				p.PlayerName, p.TileName, p.Price,
			))

			dialog.ShowConfirm("Buy Property", fmt.Sprintf("%s, do you want to buy %s for $%d?", p.PlayerName, p.TileName, p.Price), func(b bool) {
				if b {
					commandChan <- game.GameCommand{
						Type:     game.CmdBuyProperty,
						TileName: p.TileName,
					}
				}
			}, mainWindow)
		})
	})

	bus.Subscribe(events.BoughtProperty, func(g events.GameEvent) {
		fyne.Do(func() {
			p := g.Payload.(events.BoughtPropertyPayload)

			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s bought %s for $%d\n", p.PlayerName, p.TileName, p.Price,
			))
		})
	})
}
