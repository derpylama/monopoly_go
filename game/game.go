package game

import (
	"fmt"
	"math"
	"monopoly/board"
	"monopoly/common"
	"monopoly/dice"
	"monopoly/events"
	"monopoly/player"
	"monopoly/tile"
)

type Game struct {
	players       []*player.Player
	currentPlayer int
	board         *board.Board
	gameOver      bool
	dice          *dice.Dice

	//contains all events that occur during the game
	bus *common.Bus
}

func NewGame(bus *common.Bus) Game {
	// When a new game is created all players, board and dice should be initialized

	players := []*player.Player{player.NewPlayer(1500, "Player 1"), player.NewPlayer(1500, "Player 2")}

	board := board.NewBoard()
	dice := dice.NewDice(2, 6)

	return Game{
		players:       players,
		currentPlayer: 0,
		board:         board,
		gameOver:      false,
		dice:          dice,
		bus:           bus,
	}
}

func (game *Game) StartGame() {
	// Start the first turn by publishing StartTurn for player 0
	game.bus.Publish(common.GameEvent{
		Type: common.StartTurn,
		Payload: events.StartTurnPayload{
			PlayerName: game.players[game.currentPlayer].GetName(),
			Money:      game.getPlayer().GetMoney(),
			TileName:   "Go",
		},
	})

}

func (game *Game) EndGame() {

}

func (game *Game) AddPlayer() {
	//Gets all required inputs and creates a new player

}

func (game *Game) takeTurn() {
	currentPlayer := game.getPlayer()

	game.bus.Publish(common.GameEvent{
		Type: common.StartTurn,
		Payload: events.StartTurnPayload{
			PlayerName:      currentPlayer.GetName(),
			Money:           currentPlayer.GetMoney(),
			OwnedProperties: GetPlayersProperties(currentPlayer, game.board.Tiles()),
			TileName:        game.board.GetTile(currentPlayer.GetPosition()).GetName(),
		},
	})

	game.bus.Publish(common.GameEvent{
		Type: common.InputPromptOptions,
		Payload: events.InputPromptPayload{
			PlayerName: currentPlayer.GetName(),
			Options:    GetPlayerOptions(currentPlayer, "go"),
		},
	})
}

func (game *Game) nextTurn() {
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
	game.takeTurn()
}

// Returns a pointer to the current player object instance
func (game *Game) getPlayer() *player.Player {
	currentPlayerIndex := game.getCurrentPlayerIndex()
	currentPlayer := game.players[currentPlayerIndex]

	return currentPlayer
}

func (game *Game) getCurrentPlayerIndex() int {
	return game.currentPlayer
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// GetPlayersProperties returns all tiles owned by a player.
// Works for streets, utilities, railroads, or any PropertyTile-based tile.
func GetPlayersProperties(player *player.Player, tiles []common.Tile) []common.Tile {
	var properties []common.Tile

	for _, t := range tiles {
		// Attempt to extract an owned-by property using common PropertyTile methods
		switch tile := t.(type) {
		case *tile.Street:
			if tile.GetOwner() == player {
				properties = append(properties, t)
			}
		case *tile.Utility:
			if tile.GetOwner() == player {
				properties = append(properties, t)
			}
		case *tile.TrainStation:
			if tile.GetOwner() == player {
				properties = append(properties, t)
			}

		}
	}

	return properties
}

func (game *Game) Handle(cmd GameCommand) {
	player := game.getPlayer()

	println("Handling command:", cmd.Type, "for player:", player.GetName())

	switch cmd.Type {

	case CmdRollDice:
		game.handleRollDice(player)

	case CmdEndTurn:
		game.nextTurn()

	case CmdBuyProperty:
		fmt.Println("player: " + player.GetName() + "bought " + cmd.TileName)
		game.handleBuyProperty(player, cmd.TileName)

	case CmdPayToExitJail:
		game.handlePayToExitJail(player, cmd.TileName)

	case CmdMortgage:
		game.handleMortgage(player, cmd.TileName)

	case CmdUnMortgage:
		game.handleMortgage(player, cmd.TileName)
	default:
		println("Unknown command type:", string(cmd.Type))

	}

}

func (game *Game) handleRollDice(player *player.Player) {
	roll := game.dice.ThrowDice()

	if !player.GetJailStatus() {

		if !player.HasRolled() {

			player.Move(roll)

			game.bus.Publish(common.GameEvent{
				Type: common.RolledDice,
				Payload: events.RolledDicePayload{
					PlayerName: player.GetName(),
					Dice:       roll,
				},
			})

			landedOnTile := game.board.GetTile(player.GetPosition())

			game.bus.Publish(common.GameEvent{
				Type: common.LandedOnTile,
				Payload: events.LandedOnTilePayload{
					PlayerName: player.GetName(),
					TileName:   landedOnTile.GetName(),
				},
			})

			landedOnTile.OnLand(player, game.board.Tiles(), roll, game.Bus())

		}

		if diceDubbleRollCheck(roll) && player.GetAmountOfDubbles() < 3 {
			player.IncrementAmountOfDubbles()

			game.bus.Publish(common.GameEvent{
				Type: common.InputPromptOptions,
				Payload: events.InputPromptPayload{
					PlayerName: player.GetName(),
					Options:    GetPlayerOptions(player, game.board.GetTile(game.getPlayer().GetPosition()).GetName()),
				},
			})
			return
		} else {
			player.ToggleHasRolled()

			game.bus.Publish(common.GameEvent{
				Type: common.InputPromptOptions,
				Payload: events.InputPromptPayload{
					PlayerName: player.GetName(),
					Options:    GetPlayerOptions(player, game.board.GetTile(game.getPlayer().GetPosition()).GetName()),
				},
			})
		}

	} else {
		if player.GetJailedTurns() < 3 && diceDubbleRollCheck(roll) {

			player.ToggleIsJailed()
			player.ResetJailedTurns()
			player.Move(roll)

		} else if player.GetJailedTurns() == 3 {
			game.bus.Publish(common.GameEvent{
				Type: common.ForcedPayToExitJail,
				Payload: events.ForcedPayToExitJailPayload{
					PlayerName: player.GetName(),
					Price:      50,
				},
			})
		} else {

			game.bus.Publish(common.GameEvent{
				Type: common.Jailed,
				Payload: events.JailedPayload{
					PlayerName:  player.GetName(),
					JailedTurns: player.GetJailedTurns(),
				},
			})
		}

	}
}

func (game *Game) handleBuyProperty(player *player.Player, tileName string) {
	boughtTile, ok := game.board.GetTileByName(tileName)

	if ok {
		property := boughtTile.(tile.Property)
		property.BuyProperty(player, game.bus)

		game.bus.Publish(common.GameEvent{
			Type: common.UpdateMoney,
			Payload: events.UpdateMoneyPayload{
				PlayerName: player.GetName(),
				Money:      player.GetMoney(),
			},
		})

		game.bus.Publish(common.GameEvent{
			Type: common.UpdateProperties,
			Payload: events.UpdatePropertiesPayload{
				PlayerName:      player.GetName(),
				OwnedProperties: GetPlayersProperties(player, game.board.Tiles()),
			},
		})

	} else {
		game.bus.Publish(common.GameEvent{
			Type: common.InputPromptOptions,
			Payload: events.InputPromptPayload{
				PlayerName: player.GetName(),
				Options:    GetPlayerOptions(player, tileName),
			},
		})
	}
}

func (game *Game) handlePayToExitJail(player *player.Player, tileName string) {
	player.Pay(50)
	player.ToggleIsJailed()
	player.ResetJailedTurns()
}

func (game *Game) handleMortgage(player *player.Player, tileName string) {
	mortgegedTile, ok := game.board.GetTileByName(tileName)

	if ok {
		property := mortgegedTile.(tile.Property)

		if !property.GetMortgageStatus() {
			property.Mortgage()

			game.bus.Publish(common.GameEvent{
				Type: common.UpdateProperties,
				Payload: events.UpdatePropertiesPayload{
					PlayerName:      player.GetName(),
					OwnedProperties: GetPlayersProperties(player, game.board.Tiles()),
				},
			})

			game.bus.Publish(common.GameEvent{
				Type: common.UpdateMoney,
				Payload: events.UpdateMoneyPayload{
					PlayerName: player.GetName(),
					Money:      player.GetMoney(),
				},
			})

			game.bus.Publish(common.GameEvent{
				Type: common.MortgageProperty,
				Payload: events.MortgagePropertyPayload{
					PlayerName:    player.GetName(),
					TileName:      property.GetName(),
					MortgageValue: property.GetMortgageValue(),
				},
			})

		} else {
			property.UnMortgage()

			game.bus.Publish(common.GameEvent{
				Type: common.UpdateProperties,
				Payload: events.UpdatePropertiesPayload{
					PlayerName:      player.GetName(),
					OwnedProperties: GetPlayersProperties(player, game.board.Tiles()),
				},
			})

			game.bus.Publish(common.GameEvent{
				Type: common.UpdateMoney,
				Payload: events.UpdateMoneyPayload{
					PlayerName: player.GetName(),
					Money:      player.GetMoney(),
				},
			})

			game.bus.Publish(common.GameEvent{
				Type: common.MortgageProperty,
				Payload: events.MortgagePropertyPayload{
					PlayerName:    player.GetName(),
					TileName:      property.GetName(),
					MortgageValue: int(math.Round(float64(property.GetMortgageValue()) * float64(1.1))),
				},
			})
		}

	}
}

func (g *Game) Bus() *common.Bus {
	return g.bus
}

func PlayerOwnsColorSet(player *player.Player, color tile.Color, board *board.Board) bool {
	for _, t := range board.Tiles() {
		street, ok := t.(*tile.Street)
		if !ok {
			continue
		}

		if street.GetColor() == color {
			if street.GetOwner() != player {
				return false
			}
		}
	}

	return true
}

func GetPlayerOptions(player *player.Player, tileName string) []any {
	commandList := []any{}

	if !player.HasRolled() {
		commandList = append(commandList, GameCommand{Type: CmdRollDice, PlayerName: player.GetName(), TileName: tileName})
	}

	if player.GetJailStatus() {
		commandList = append(commandList, GameCommand{Type: CmdPayToExitJail, PlayerName: player.GetName(), TileName: tileName})
	}

	commandList = append(commandList, GameCommand{Type: CmdEndTurn, PlayerName: player.GetName(), TileName: tileName})

	return commandList
}

func diceDubbleRollCheck(dice []int) bool {
	return dice[0] == dice[1]
}
