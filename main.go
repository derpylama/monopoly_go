package main

import (
	"monopoly/board"
	"monopoly/dice"
	"monopoly/game"
	"monopoly/player"
)

func main() {

	players := []*player.Player{player.NewPlayer(1500, "janne")}
	board := board.NewBoard()
	dice := dice.NewDice(2, 6)

	game := game.NewGame(players, board, dice)

	game.StartGame()
}
