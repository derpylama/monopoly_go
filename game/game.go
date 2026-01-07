package game

import (
	"bufio"
	"fmt"
	"monopoly/board"
	"monopoly/dice"
	"monopoly/logger"
	"monopoly/player"
	"monopoly/tile"

	"os"
	"strconv"
	"strings"
)

type Game struct {
	players       []*player.Player
	currentPlayer int
	board         *board.Board
	gameOver      bool
	dice          *dice.Dice
}

func NewGame(players []*player.Player, board *board.Board, dice *dice.Dice) Game {
	return Game{
		players:       players,
		currentPlayer: 0,
		board:         board,
		gameOver:      false,
		dice:          dice,
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

func (game *Game) takeTurn() {
	currentPlayer := game.players[game.currentPlayer]
	var reader = bufio.NewReader(os.Stdin)

	fmt.Print(currentPlayer.GetName() + " turn")
	fmt.Print("\nEnter r to roll dice: \n\n")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "r":
		roll := game.dice.ThrowDice()

		logger.LogRollDice(currentPlayer.GetName(), roll, currentPlayer.GetMoney())

		currentPos := currentPlayer.GetPosition()

		currentPlayer.Move(roll)
		landedOnTile := game.board.GetTile(currentPlayer.GetPosition())

		allTiles := game.board.Tiles()

		if currentPlayer.GetPosition() < currentPos {
			//Passed GO
			currentPlayer.SetMoney(currentPlayer.GetMoney() + 200)
		}

		switch v := landedOnTile.(type) {
		case *tile.Street:

			if v.IsOwned() {

				//If the property is owned then get the rent and pay it
				logger.LogOnLand(v.GetOwner().GetName(), v.GetName(), true, v.GetRent(allTiles, roll), currentPlayer.GetName(), v)
				playerPaysRent(currentPlayer, v.GetRent(allTiles, roll), v.GetOwner())

				playerInputs(currentPlayer, allTiles)
			} else {

				//If the property is not owned get the price of it
				logger.LogOnLand("", v.GetName(), false, 0, currentPlayer.GetName(), v)

				// Ask the player if they want to buy the property
				fmt.Print("Do you want to buy " + v.GetName() + " for " + strconv.Itoa(v.GetPrice()) + "? (y/n): ")
				buyInput, _ := reader.ReadString('\n')
				buyInput = strings.TrimSpace(buyInput)

				if buyInput == "y" {
					playerBuysProperty(currentPlayer, v)
					remainingMoney := currentPlayer.GetMoney()

					logger.LogBuyProperty(currentPlayer.GetName(), v.GetName(), v.GetPrice(), remainingMoney)

					playerInputs(currentPlayer, allTiles)
					// fmt.Print(currentPlayer.GetName() + " bought " + v.GetName() + " for " + strconv.Itoa(v.GetPrice()) + ". Remaining money: " + strconv.Itoa(remainingMoney) + "\n")
				} else {
					fmt.Print("You chose not to buy the property.\n")
				}
			}

		case *tile.TrainStation:
			if v.IsOwned() {

				logger.LogOnLand(v.GetOwner().GetName(), v.GetName(), true, v.GetRent(allTiles, roll), currentPlayer.GetName(), v)
				playerPaysRent(currentPlayer, v.GetRent(allTiles, roll), v.GetOwner())
			}

		case *tile.TaxTile:
			logger.LogOnLand("", v.GetName(), false, v.GetTaxAmount(), currentPlayer.GetName(), v)
			v.OnLand(currentPlayer)

			playerInputs(currentPlayer, allTiles)

		case *tile.Utility:
			if v.IsOwned() {

				logger.LogOnLand(v.GetOwner().GetName(), v.GetName(), true, v.GetRent(allTiles, roll), currentPlayer.GetName(), v)
				playerPaysRent(currentPlayer, v.GetRent(allTiles, roll), v.GetOwner())
			} else {
				logger.LogOnLand("", v.GetName(), false, 0, currentPlayer.GetName(), v)

			}

		default:
			//if not implemented then just log the tile landing
			logger.LogOnLand("", v.GetName(), false, 0, currentPlayer.GetName(), v)

		}

	case "a":

	default:
		fmt.Println("Invalid Key press")
	}

	ClearScreen()
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

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
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

func playerInputs(player *player.Player, tiles []tile.Tile) {
	// checks what input the player gives

	fmt.Println("Press 'a' to view all your properties or b to buy a house any other key to continue:")

	var inputReader = bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "a":
		logger.LogPlayersProperties(player.GetName(), GetPlayersProperties(player, tiles))

	case "b":
		// buy houses or hotels
		logger.LogPlayersProperties(player.GetName(), GetPlayersProperties(player, tiles))

		fmt.Println("Enter the numbet of the property that you want to build a house on")
		houseInput, _ := inputReader.ReadString('\n')
		houseInput = strings.TrimSpace(houseInput)
		houseIndex, err := strconv.Atoi(houseInput)

		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		properties := GetPlayersProperties(player, tiles)
		if houseIndex < 0 || houseIndex >= len(properties) {
			fmt.Println("Invalid property index")
			return
		}

		selectedProperty := properties[houseIndex]

		switch v := selectedProperty.(type) {
		case *tile.Street:
			v.BuyHouse()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			player.Pay(v.GetHousePrice())
			fmt.Println("Built a house on " + v.GetName())
			playerInputs(player, tiles)
		default:
			playerInputs(player, tiles)
			return
		}

	}
}

func GetPlayersProperties(player *player.Player, tiles []tile.Tile) []tile.Property {
	var properties []tile.Property

	for _, t := range tiles {
		// assert pointer type because PropertyTile's methods have pointer receivers
		pt, ok := t.(tile.Property)
		if !ok {
			continue
		}

		if pt.GetOwner() != nil && pt.GetOwner().GetName() == player.GetName() {
			properties = append(properties, pt)
		}
	}

	return properties
}
