package main

import (
	"flag"
	"fmt"
	"monopoly/common"
	"monopoly/game"
	"monopoly/logger"
	"monopoly/view"
)

func main() {
	ClearTerminal()
	uiPtr := flag.String("ui", "console", "Type of UI to use: console or gui")
	flag.Parse()

	if *uiPtr != "console" && *uiPtr != "gui" {
		panic("Invalid UI type. Use 'console' or 'gui'.")
	}

	bus := common.NewBus()
	commandChannel := make(chan game.GameCommand, 10)

	g := game.NewGame(bus)

	g.RegisterListeners(commandChannel)
	g.RegisterPromptListeners(commandChannel)

	if *uiPtr == "console" {
		log := logger.New(true)

		logger.RegisterListeners(bus, log)
		logger.RegisterPromptListeners(bus, commandChannel)

		// Start command proccesing goroutine
		go func() {
			for cmd := range commandChannel {
				g.Handle(cmd)
			}
		}()
		g.StartGame()

	}

	view.StartGUI(&g, bus, commandChannel)
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
