package main

import (
	"monopoly/board"
	"monopoly/dice"
	"monopoly/game"
	"monopoly/player"

	"fyne.io/fyne/v2/app"
)

func main() {

	players := []*player.Player{player.NewPlayer(1500, "player1"), player.NewPlayer(1500, "player2")}
	board := board.NewBoard()
	dice := dice.NewDice(2, 6)

	game := game.NewGame()

	go game.StartGame()

	mainContext := app.New()
	mainWindow := mainContext.NewWindow("Monopoly")
	mainWindow.ShowAndRun()

}
