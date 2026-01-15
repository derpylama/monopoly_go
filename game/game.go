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
	// game.takeTurn()
	// Start the first turn by publishing StartTurn for player 0
	game.bus.Publish(common.GameEvent{
		Type: common.StartTurn,
		Payload: events.StartTurnPayload{
			PlayerName: game.players[game.currentPlayer].GetName(),
			Money:      game.getPlayer().GetMoney(),
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
		},
	})

	game.bus.Publish(common.GameEvent{
		Type: common.InputPromptOptions,
		Payload: events.InputPromptPayload{
			PlayerName: currentPlayer.GetName(),
			Options:    []any{GameCommand{Type: CmdEndTurn, PlayerName: currentPlayer.GetName()}, GameCommand{Type: CmdRollDice, PlayerName: currentPlayer.GetName()}},
		},
	})
}

func (game *Game) nextTurn() {
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
	game.takeTurn()
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

	println("Handling command:", cmd.Type, "for player:", player.GetName())

	switch cmd.Type {

	case CmdRollDice:
		game.handleRollDice(player)

	case CmdEndTurn:
		game.nextTurn()

	case CmdBuyProperty:
		fmt.Println("player: " + player.GetName() + "bought " + cmd.TileName)
		game.handleBuyProperty(player, cmd.TileName)

	default:
		println("Unknown command type:", string(cmd.Type))

	}

}

func (game *Game) handleRollDice(player *player.Player) {
	roll := game.dice.ThrowDice()
	player.Move(roll)

	//pos, _ := common.GetTilePosByName("Jail", game.board.Tiles())

	//player.Teleport(pos)

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

	game.bus.Publish(common.GameEvent{
		Type: common.InputPromptOptions,
		Payload: events.InputPromptPayload{
			PlayerName: player.GetName(),
			Options:    []any{GameCommand{Type: CmdEndTurn, PlayerName: player.GetName()}, GameCommand{Type: CmdRollDice, PlayerName: player.GetName()}},
		},
	})
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
	} else {
		game.bus.Publish(common.GameEvent{
			Type: common.InputPromptOptions,
			Payload: events.InputPromptPayload{
				PlayerName: player.GetName(),
				Options:    []any{GameCommand{Type: CmdEndTurn, PlayerName: player.GetName()}, GameCommand{Type: CmdRollDice, PlayerName: player.GetName()}},
			},
		})
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

func GetPlayerOptions(player *player.Player) {

}
