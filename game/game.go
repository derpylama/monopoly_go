package game

import (
	"bufio"
	"fmt"
	"monopoly/board"
	"monopoly/dice"
	"monopoly/helper"
	"monopoly/player"
	"monopoly/tile"

	"os"
	"strconv"
	"strings"
)

type Game struct {
	players       []*player.Player
	currentPlayer int
	board         board.Board
	gameOver      bool
	dice          *dice.Dice
}

func NewGame(players []*player.Player, board *board.Board, dice *dice.Dice) Game {
	return Game{
		players:       players,
		currentPlayer: 0,
		board:         *board,
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
		fmt.Println("Player " + currentPlayer.GetName() + " rolled a " + strconv.Itoa(helper.SumOfList(roll)))
		currentPlayer.Move(roll)
		landedOnTile := game.board.GetTile(currentPlayer.GetPosition())

		fmt.Println(currentPlayer.GetName() + " is now on " + landedOnTile.GetName() + ". Position " + strconv.Itoa(currentPlayer.GetPosition()))

		switch v := landedOnTile.(type) {
		case *tile.Street:
			if v.GetOwner() != nil {
				fmt.Println("The Property is owned by " + v.GetOwner().GetName())
			} else {
				fmt.Print("Property is not owned\n")
				fmt.Print("Price: " + strconv.Itoa(v.GetPrice()) + "\n")
			}
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

func playerBuysProperty() {

}
