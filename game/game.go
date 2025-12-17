package game

import (
	"fmt"
	"monopoly/board"
	"monopoly/dice"
	"monopoly/helper"
	"monopoly/player"
	"strconv"
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

	var input string

	fmt.Println("Enter r to roll dice: \nEnter a to ")
	fmt.Scanln(&input)
	ClearScreen()

	switch input {
	case "r":
		roll := game.dice.ThrowDice()
		fmt.Println("Player " + currentPlayer.GetName() + " rolled " + strconv.Itoa(helper.SumOfList(roll)))
		currentPlayer.Move(roll)
		tile := game.board.GetTile(currentPlayer.GetPosition())

		fmt.Println(currentPlayer.GetName() + " is now on " + tile.GetName())

	case "a":

	default:
		fmt.Println("Invalid Key press")
	}

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
