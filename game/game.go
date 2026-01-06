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
		// fmt.Println("Player " + currentPlayer.GetName() + " rolled a " + strconv.Itoa(helper.SumOfList(roll)))
		// fmt.Print("\nCurrent money: " + strconv.Itoa(currentPlayer.GetMoney()) + "\n")

		logger.LogRollDice(currentPlayer.GetName(), roll, currentPlayer.GetMoney())

		currentPlayer.Move(roll)
		landedOnTile := game.board.GetTile(currentPlayer.GetPosition())

		allTiles := game.board.Tiles()

		// fmt.Println(currentPlayer.GetName() + " is now on " + landedOnTile.GetName() + ". Position " + strconv.Itoa(currentPlayer.GetPosition()))
		//logger.LogOnLand("", landedOnTile.GetName(), false, 0, currentPlayer.GetName(), landedOnTile)

		switch v := landedOnTile.(type) {
		case *tile.Street:

			if v.IsOwned() {

				//If the property is owned then get the rent and pay it

				logger.LogOnLand(v.GetOwner().GetName(), v.GetName(), true, v.GetRent(allTiles, roll), currentPlayer.GetName(), v)
				playerPaysRent(currentPlayer, v.GetRent(allTiles, roll), v.GetOwner())
			} else {

				//If the property is not owned get the price of it
				logger.LogOnLand("", v.GetName(), false, 0, currentPlayer.GetName(), v)

				// Ask the player if they want to buy the property
				fmt.Print("Do you want to buy " + v.GetName() + " for " + strconv.Itoa(v.GetPrice()) + "? (y/n): ")
				buyInput, _ := reader.ReadString('\n')
				buyInput = strings.TrimSpace(buyInput)

				if buyInput == "y" {
					//previousMoney := currentPlayer.GetMoney()
					playerBuysProperty(currentPlayer, v)
					remainingMoney := currentPlayer.GetMoney()

					logger.LogBuyProperty(currentPlayer.GetName(), v.GetName(), v.GetPrice(), remainingMoney)

					// fmt.Print(currentPlayer.GetName() + " bought " + v.GetName() + " for " + strconv.Itoa(v.GetPrice()) + ". Remaining money: " + strconv.Itoa(remainingMoney) + "\n")
				} else {
					fmt.Print("You chose not to buy the property.\n")
				}
			}

		case *tile.TrainStation:
			if v.IsOwned() {
				// fmt.Print("The Property is owned by " + v.GetOwner().GetName())
				// fmt.Print("\nYou must pay " + strconv.Itoa(v.GetRent(allTiles, roll)) + "\n")

				// fmt.Print("\nYour new balance is " + strconv.Itoa(currentPlayer.GetMoney()))

				logger.LogOnLand(v.GetOwner().GetName(), v.GetName(), true, v.GetRent(allTiles, roll), currentPlayer.GetName(), v)
			}

		case *tile.TaxTile:
			fmt.Print("\n Must pay " + strconv.Itoa(v.GetTaxAmount()) + "\n")
			v.OnLand(currentPlayer)
			fmt.Print("Current money: " + strconv.Itoa(currentPlayer.GetMoney()))

		case *tile.Utility:

		}

	case "a":

	default:
		fmt.Println("Invalid Key press")
	}

	fmt.Println("Press enter to continue to next turn")
	reader.ReadString('\n')

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
		fmt.Print("\nYou can't afford this property")
	}
}

func playerPaysRent(player *player.Player, amount int, owner *player.Player) {
	if player.Pay(amount) {
		owner.SetMoney(owner.GetMoney() + amount)
	} else {
		fmt.Print("\nYou can't afford to pay the rent")
	}
}
