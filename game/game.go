package game

import (
	"fmt"
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
	bus *events.Bus
}

func NewGame(bus *events.Bus) Game {
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
	for !game.gameOver {
		game.takeTurn()
		game.nextPlayer()
	}

}

func (game *Game) EndGame() {

}

func (game *Game) AddPlayer() {
	//Gets all required inputs and creates a new player

}

func (game *Game) takeTurn() {
	currentPlayer := game.players[game.currentPlayer]

	game.bus.Publish(events.GameEvent{
		Type: events.InputPromptRollDice,
		Payload: events.RolledDicePayload{
			PlayerName: currentPlayer.GetName(),
		},
	})

	//logger.LogOnLandInfo(events.GameEvent{PlayerName: currentPlayer.GetName(), TileName: landedOnTile.GetName(), Details: currentPlayer.GetName() + " landed on " + landedOnTile.GetName()})

	// for !turnIsOver {
	// 	input := inputhandler.PlayerTurnInteraction(currentPlayer.GetName())
	// 	switch input {
	// 	case "end":
	// 		if hasRolled {
	// 			turnIsOver = true

	// 		} else {
	// 			fmt.Println("You must roll the dice before ending your turn.")
	// 		}

	// 	case "trade":
	// 		// Trade logic here
	// 		fmt.Println("Trade feature is not implemented yet.")

	// 	case "mortgage":
	// 		// Mortgage logic here
	// 		fmt.Println("Mortgage feature is not implemented yet.")

	// 	case "unmortgage":
	// 		// Unmortgage logic here
	// 		fmt.Println("Unmortgage feature is not implemented yet.")

	// 	case "build":
	// 		// Build logic here
	// 		properties := GetPlayersProperties(currentPlayer, game.board.Tiles())

	// 		var propertyNames []string

	// 		for _, property := range properties {
	// 			propertyNames = append(propertyNames, property.GetName())
	// 		}

	// 		if len(properties) == 0 {
	// 			fmt.Println("You do not own any properties to build on.")
	// 		} else {
	// 			logger.LogOwnedProperties(currentPlayer.GetName(), propertyNames)

	// 			streetToBuildOn := inputhandler.PlayerEnterNumber("Enter the number of the property you want to build a house on:")

	// 			street, ok := properties[streetToBuildOn].(*tile.Street)
	// 			if !ok {
	// 				fmt.Println("You can only build houses on street properties.")
	// 				continue
	// 			}
	// 			if inputhandler.PlayerWantsToBuildHouse(currentPlayer.GetName(), street.GetName(), street.GetHousePrice()) {
	// 				if currentPlayer.Pay(street.GetHousePrice()) {
	// 					street.BuyHouse()
	// 					fmt.Println("Built a house on", street.GetName())
	// 				} else {
	// 					fmt.Println("You cannot afford to build a house on", street.GetName())
	// 				}
	// 			}
	// 		}

	// 	case "list":
	// 		properties := GetPlayersProperties(currentPlayer, game.board.Tiles())
	// 		var propertyNames []string

	// 		for _, property := range properties {
	// 			propertyNames = append(propertyNames, property.GetName())
	// 		}
	// 		logger.LogOwnedProperties(currentPlayer.GetName(), propertyNames)
	// 	case "roll":
	// 		if !hasRolled {
	// 			roll := game.dice.ThrowDice()
	// 			currentPlayer.Move(roll)
	// 			hasRolled = true

	// 			landedOnTile = game.board.GetTile(currentPlayer.GetPosition())

	// 			//logger.LogOnLandInfo(events.GameEvent{PlayerName: currentPlayer.GetName(), TileName: landedOnTile.GetName(), Details: currentPlayer.GetName() + " landed on " + landedOnTile.GetName()})
	// 			eventList := landedOnTile.OnLand(currentPlayer, game.board.Tiles(), roll)
	// 			logger.LogEvent(eventList)
	// 		} else if hasRolled {
	// 			fmt.Println("You have already rolled the dice this turn.")
	// 		} else {

	// 			fmt.Println("You are in Jail! You cannot roll the dice.")
	// 		}
	// 	default:
	// 		fmt.Println("Invalid input, please try again.")
	// 	}
	// }

}

func (game *Game) nextPlayer() {
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}

func (game *Game) getPlayer() *player.Player {
	currentPlayerIndex := game.getCurrentPlayerIndex()
	currentPlayer := game.players[currentPlayerIndex]

	return currentPlayer
}

func (game *Game) getCurrentPlayerIndex() int {
	return game.currentPlayer
}

func playerBuysProperty(player *player.Player, tile tile.Property) {
	if player.Pay(tile.GetPrice()) {
		tile.SetOwner(player)
	} else {
		fmt.Println("You can't afford this property")
	}
}

func playerPaysRent(player *player.Player, amount int, owner *player.Player) {
	if player.Pay(amount) {
		owner.SetMoney(owner.GetMoney() + amount)
	} else {
		fmt.Println("You can't afford to pay the rent")
	}
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

	switch cmd.Type {

	case CmdRollDice:
		game.handleRollDice(player)
	}

}

func (game *Game) handleRollDice(player *player.Player) {
	roll := game.dice.ThrowDice()
	player.Move(roll)

	game.bus.Publish(events.GameEvent{
		Type: events.RolledDice,
		Payload: events.RolledDicePayload{
			PlayerName: player.GetName(),
			Dice:       roll,
		},
	})

	landedOnTile := game.board.GetTile(player.GetPosition())

	game.bus.Publish(events.GameEvent{
		Type: events.LandedOnTile,
		Payload: events.LandedOnTilePayload{
			PlayerName: player.GetName(),
			TileName:   landedOnTile.GetName(),
		},
	})

	landedOnTile.OnLand(player, game.board.Tiles(), roll, game.Bus())
}

func (g *Game) Bus() *events.Bus {
	return g.bus
}
