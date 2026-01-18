package view

import (
	"fmt"
	"image/color"
	"monopoly/common"
	"monopoly/events"
	"monopoly/game"
	"monopoly/tile"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func appendLogText(logArea *widget.Entry, newText string, maxLines int) {
	currentText := logArea.Text
	lines := strings.Split(currentText, "\n")
	if len(lines) >= maxLines {
		lines = lines[len(lines)-maxLines+1:]
	}
	newFullText := strings.Join(lines, "\n") + newText
	logArea.SetText(newFullText)
}

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

			LoadOwnedProperties(propertiesCon, payload.OwnedProperties, commandChan, payload.AllTiles)
			propertiesCon.Refresh()
		})
	})

	bus.Subscribe(common.UpdateProperties, func(ge common.GameEvent) {
		payload := ge.Payload.(events.UpdatePropertiesPayload)

		fyne.Do(func() {
			LoadOwnedProperties(propertiesCon, payload.OwnedProperties, commandChan, payload.AllTiles)
		})
	})

	bus.Subscribe(common.Jailed, func(ge common.GameEvent) {
		payload := ge.Payload.(events.JailedPayload)

		fyne.Do(func() {
			appendLogText(logArea,
				fmt.Sprintf("%s is in jail and has been for %d turns\n",
					payload.PlayerName, payload.JailedTurns),
				50,
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

			payToExitJailDialog.Show()
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

		fyne.Do(func() {
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

			buttonContainer.Refresh()
		})

	})

	// Logging events
	bus.Subscribe(common.RolledDice, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.RolledDicePayload)
			appendLogText(logArea, fmt.Sprintf(
				"%s rolled %d + %d = %d\n",
				p.PlayerName, p.Dice[0], p.Dice[1], p.Dice[0]+p.Dice[1],
			), 50)
		})
	})

	bus.Subscribe(common.LandedOnTile, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.LandedOnTilePayload)
			appendLogText(logArea, fmt.Sprintf(
				"%s landed on %s\n",
				p.PlayerName, p.TileName,
			), 50)
		})
	})

	bus.Subscribe(common.PaidRent, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.PaidRentPayload)
			appendLogText(logArea, fmt.Sprintf(
				"%s paid $%d rent to %s for %s\n",
				p.PlayerName, p.Rent, p.Owner, p.TileName,
			), 50)

		})
	})

	bus.Subscribe(common.PaidTax, func(e common.GameEvent) {
		fyne.Do(func() {
			p := e.Payload.(events.TaxPayload)
			appendLogText(logArea, fmt.Sprintf(
				"%s paid $%d tax at %s\n",
				p.PlayerName, p.TaxAmount, p.TileName,
			), 50)

		})
	})

	bus.Subscribe(common.LandedOnUnownedProperty, func(ge common.GameEvent) {
		fyne.Do(func() {
			p := ge.Payload.(events.LandedOnUnownedPropertyPayload)
			appendLogText(logArea, fmt.Sprintf(
				"%s landed on unowned property %s (cost: $%d)\n",
				p.PlayerName, p.TileName, p.Price,
			), 50)

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

	bus.Subscribe(common.MortgageProperty, func(ge common.GameEvent) {
		fyne.Do(func() {
			p := ge.Payload.(events.MortgagePropertyPayload)

			appendLogText(logArea, fmt.Sprintf(
				"%s mortgaged %s for %d\n", p.PlayerName, p.TileName, p.MortgageValue,
			), 50)
		})
	})

	bus.Subscribe(common.UnMortgageProperty, func(ge common.GameEvent) {
		fyne.Do(func() {
			p := ge.Payload.(events.MortgagePropertyPayload)

			appendLogText(logArea, fmt.Sprintf(
				"%s un morgtgage %s and paid %d\n", p.PlayerName, p.TileName, p.MortgageValue,
			), 50)
		})
	})

	bus.Subscribe(common.BuiltHouse, func(ge common.GameEvent) {
		p := ge.Payload.(events.BuiltHousePayload)

		fyne.Do(func() {
			appendLogText(logArea, fmt.Sprintf(
				"%s built a house on %s for $%d and now has $%d in rent\n", p.PlayerName, p.TileName, p.HousePrice, p.NewRent,
			), 50)
		})

	})

	bus.Subscribe(common.BuiltHotel, func(ge common.GameEvent) {
		p := ge.Payload.(events.BuiltHotelPayload)

		fyne.Do(func() {
			appendLogText(logArea, fmt.Sprintf(
				"%s built a Hotel on %s for $%d and now has $%d in rent\n", p.PlayerName, p.TileName, p.HotelPrice, p.NewRent,
			), 50)
		})
	})

	bus.Subscribe(common.OwnMaxAmountOfHouses, func(ge common.GameEvent) {
		p := ge.Payload.(events.OwnMaxAmountOfHousesPayload)

		fyne.Do(func() {
			appendLogText(logArea, fmt.Sprintf(
				"%s owns the max amount of houses on %s\n", p.PlayerName, p.TileName,
			), 50)
		})
	})

	bus.Subscribe(common.BoughtProperty, func(g common.GameEvent) {
		fyne.Do(func() {
			p := g.Payload.(events.BoughtPropertyPayload)

			appendLogText(logArea, fmt.Sprintf(
				"%s bought %s for $%d\n", p.PlayerName, p.TileName, p.Price,
			), 50)
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

func LoadOwnedProperties(propertiesCon *fyne.Container, ownedProperties []common.Tile, commandChan chan<- game.GameCommand, allTiles []common.Tile) {
	fyne.Do(func() {

		propertiesCon.Objects = nil

		if len(ownedProperties) > 0 {
			for _, property := range ownedProperties {
				property := property.(tile.Property)

				propLabel := widget.NewLabel(fmt.Sprintf("%s Rent: $%d", property.GetName(), property.GetRent(allTiles, []int{0, 1})))

				propertiesCon.Add(propLabel)

				switch property := property.(type) {
				case *tile.Street:

					if property.GetMortgageStatus() {
						mortgageBtn := widget.NewButton("UnMortgage", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdMortgage,
								TileName: property.GetName(),
							}
						})

						propertiesCon.Add(mortgageBtn)
					} else {
						mortgageBtn := widget.NewButton("Mortgage", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdMortgage,
								TileName: property.GetName(),
							}

						})

						propertiesCon.Add(mortgageBtn)
					}

					bg := canvas.NewRectangle(convertColorStringToRGB(property.GetColor()))
					bg.SetMinSize(fyne.NewSize(120, 100))

					if property.GetHouseAmount() < 3 {
						houseBtn := widget.NewButton("Build house", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdBuildHouse,
								TileName: property.GetName(),
							}
						})

						con := container.NewStack(bg, container.NewVBox(houseBtn))
						propertiesCon.Add(con)
					} else {
						houseBtn := widget.NewButton("Build hotel", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdBuildHotel,
								TileName: property.GetName(),
							}
						})

						con := container.NewStack(bg, container.NewVBox(houseBtn))
						propertiesCon.Add(con)
					}

				case *tile.TrainStation:

					if property.GetMortgageStatus() {
						mortgageBtn := widget.NewButton("UnMortgage", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdMortgage,
								TileName: property.GetName(),
							}
						})

						propertiesCon.Add(mortgageBtn)
					} else {
						mortgageBtn := widget.NewButton("Mortgage", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdMortgage,
								TileName: property.GetName(),
							}
						})

						propertiesCon.Add(mortgageBtn)
					}

				case *tile.Utility:

					if property.GetMortgageStatus() {
						mortgageBtn := widget.NewButton("UnMortgage", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdMortgage,
								TileName: property.GetName(),
							}
						})

						propertiesCon.Add(mortgageBtn)
					} else {
						mortgageBtn := widget.NewButton("Mortgage", func() {
							commandChan <- game.GameCommand{
								Type:     game.CmdMortgage,
								TileName: property.GetName(),
							}
						})

						propertiesCon.Add(mortgageBtn)
					}
				}

			}
		}
		propertiesCon.Refresh()
	})
}
