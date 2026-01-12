package main

import (
	"monopoly/events"
	"monopoly/game"
	"monopoly/logger"
)

func main() {

	// players := []*player.Player{player.NewPlayer(1500, "player1"), player.NewPlayer(1500, "player2")}
	// board := board.NewBoard()
	// dice := dice.NewDice(2, 6)

	bus := events.NewBus()
	commandChannel := make(chan game.GameCommand)
	game := game.NewGame(bus)

	// 1️⃣ Create logger
	log := logger.New(true)

	go game.StartGame()

	logger.RegisterListeners(bus, log)
	logger.RegisterPromptListeners(bus, commandChannel)

	go func() {
		for cmd := range commandChannel {
			game.Handle(cmd)
		}
	}()

	// mainContext := app.New()
	// mainWindow := mainContext.NewWindow("Monopoly")
	// mainWindow.ShowAndRun()
	game.StartGame()
}
