package view

import (
	"fmt"
	"image/color"
	"monopoly/common"
	"monopoly/events"
	"monopoly/game"
	"monopoly/tile"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func RegisterPromptListeners(commandChan chan<- game.GameCommand, bus *common.Bus) {

}

func RegisterReactiveGUIListeners(
	bus *common.Bus,
	g *game.Game,
	commandChan chan<- game.GameCommand,
	logArea *widget.Entry,
	buttonContainer *fyne.Container,
	playerLabel *widget.Label,
	propertiesCon *fyne.Container,
	mainWindow fyne.Window,
) {

	//Update which player's turn it is and their money
	bus.Subscribe(common.StartTurn, func(ge common.GameEvent) {
		payload := ge.Payload.(events.StartTurnPayload)

		fyne.Do(func() {
			playerLabel.SetText(
				fmt.Sprintf("%s's turn:\n%s has $%d", payload.PlayerName, payload.PlayerName, payload.Money),
			)

			propertiesCon.Objects = nil
			properties := payload.OwnedProperties

			if len(properties) > 0 {
				for _, property := range properties {

					street, ok := property.(*tile.Street)

					propLabel := widget.NewLabel(property.GetName())

					if ok {

						bg := canvas.NewRectangle(convertColorStringToRGB(street.GetColor()))
						bg.SetMinSize(fyne.NewSize(100, 100))

						con := container.NewStack(bg, propLabel)

						propertiesCon.Add(con)
					} else {
						con := container.NewStack(propLabel)

						propertiesCon.Add(con)
					}

				}
			}
			propertiesCon.Refresh()
		})
	})

	bus.Subscribe(common.UpdateProperties, func(ge common.GameEvent) {
		payload := ge.Payload.(events.UpdatePropertiesPayload)

		fyne.Do(func() {
			LoadOwnedProperties(propertiesCon, payload.OwnedProperties)
		})
	})

	bus.Subscribe(common.Jailed, func(ge common.GameEvent) {
		payload := ge.Payload.(events.JailedPayload)

		fyne.Do(func() {
			playerLabel.SetText(
				fmt.Sprintf(logArea.Text+"%s is in jail and has been for %d turns", payload.PlayerName, payload.JailedTurns),
			)
		})
	})

	bus.Subscribe(common.ForcedPayToExitJail, func(ge common.GameEvent) {
		payload := ge.Payload.(events.ForcedPayToExitJailPayload)

		fyne.Do(func() {
			payToExitJailDialog := dialog.NewCustom(
				"Notice",
				"Pay",
				widget.NewLabel(fmt.Sprintf("You have been in jail for three turns and need to pay $%d", payload.Price)),
				mainWindow,
			)

			payToExitJailDialog.SetOnClosed(func() {
				commandChan <- game.GameCommand{
					Type:       game.CmdPayToExitJail,
					PlayerName: payload.PlayerName,
				}
			})
		})
	})

	//Update the players money after every time they spend it
	bus.Subscribe(common.UpdateMoney, func(ge common.GameEvent) {
		payload := ge.Payload.(events.UpdateMoneyPayload)

		fyne.Do(func() {
			playerLabel.SetText(
				fmt.Sprintf("%s's turn:\n%s has $%d", payload.PlayerName, payload.PlayerName, payload.Money),
			)
		})
	})

	// Show prompt options from the game
	bus.Subscribe(common.InputPromptOptions, func(e common.GameEvent) {
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
	bus.Subscribe(common.RolledDice, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.RolledDicePayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s rolled %d + %d = %d\n",
				p.PlayerName, p.Dice[0], p.Dice[1], p.Dice[0]+p.Dice[1],
			))
		})
	})

	bus.Subscribe(common.LandedOnTile, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.LandedOnTilePayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s landed on %s\n",
				p.PlayerName, p.TileName,
			))
		})
	})

	bus.Subscribe(common.PaidRent, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.PaidRentPayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s paid $%d rent to %s for %s\n",
				p.PlayerName, p.Rent, p.Owner, p.TileName,
			))

		})
	})

	bus.Subscribe(common.PaidTax, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.TaxPayload)
			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s paid $%d tax at %s\n",
				p.PlayerName, p.TaxAmount, p.TileName,
			))

		})
	})

	bus.Subscribe(common.LandedOnUnownedProperty, func(ge common.GameEvent) {
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

	bus.Subscribe(common.BoughtProperty, func(g common.GameEvent) {
		fyne.Do(func() {
			p := g.Payload.(events.BoughtPropertyPayload)

			logArea.SetText(logArea.Text + fmt.Sprintf(
				"%s bought %s for $%d\n", p.PlayerName, p.TileName, p.Price,
			))
		})
	})

}

func convertColorStringToRGB(c tile.Color) color.RGBA {
	switch c {
	case tile.Brown:
		return color.RGBA{149, 84, 54, 255}

	case tile.LightBlue:
		return color.RGBA{170, 224, 252, 255}

	case tile.Pink:
		return color.RGBA{218, 57, 150, 255}

	case tile.Orange:
		return color.RGBA{247, 148, 29, 255}

	case tile.Red:
		return color.RGBA{237, 27, 36, 255}

	case tile.Yellow:
		return color.RGBA{254, 241, 2, 255}

	case tile.Green:
		return color.RGBA{31, 178, 90, 255}

	case tile.DarkBlue:
		return color.RGBA{1, 113, 187, 255}
	}

	return color.RGBA{}
}

func LoadOwnedProperties(propertiesCon *fyne.Container, ownedProperties []common.Tile) {
	fyne.Do(func() {

		propertiesCon.Objects = nil

		if len(ownedProperties) > 0 {
			for _, property := range ownedProperties {

				street, streetOk := property.(*tile.Street)

				propLabel := widget.NewLabel(property.GetName())

				if streetOk {

					bg := canvas.NewRectangle(convertColorStringToRGB(street.GetColor()))
					bg.SetMinSize(fyne.NewSize(100, 100))

					mortgageBtn := widget.NewButton("Mortgage", func() {})
					houseBtn := widget.NewButton("Build house", func() {})

					con := container.NewStack(bg, container.NewVBox(mortgageBtn, houseBtn))

					propertiesCon.Add(con)
				} else {
					con := container.NewStack(propLabel)

					propertiesCon.Add(con)
				}

			}
		}
		propertiesCon.Refresh()
	})
}
