package view

import (
	"monopoly/events"
	"monopoly/game"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartGUI(g *game.Game, bus *events.Bus, commandChan chan game.GameCommand) {
	appContext := app.New()
	mainWindow := appContext.NewWindow("Monopoly")

	//load house image
	monopolyHouse := canvas.NewImageFromFile("images/monopoly-house.png")
	monopolyHouse.FillMode = canvas.ImageFillContain
	monopolyHouse.Resize(fyne.NewSize(30, 30))

	//load hotel image
	monopolyHotel := canvas.NewImageFromFile("images/monopoly-hotel.png")
	monopolyHotel.FillMode = canvas.ImageFillContain
	monopolyHotel.Resize(fyne.NewSize(30, 30))
	monopolyHotel.Move(fyne.NewPos(300, 100))

	houseCon := container.NewWithoutLayout(monopolyHotel)

	//Load the game board
	monopolyBoard := canvas.NewImageFromFile("images/monopoly_board.jpg")
	monopolyBoard.FillMode = canvas.ImageFillContain
	monopolyBoard.SetMinSize(fyne.NewSize(600, 600))

	logArea := widget.NewMultiLineEntry()

	playerLabel := widget.NewLabel("player")

	buttons := container.NewVBox()

	leftPanel := container.NewVBox(
		playerLabel,
		buttons,
		logArea,
	)

	windowGrid := container.New(layout.NewHBoxLayout(), leftPanel, container.NewStack(monopolyBoard, houseCon), layout.NewSpacer())

	mainWindow.Resize(fyne.NewSize(1200, 800))
	mainWindow.SetContent(windowGrid)

	// Run the GUI

	// GUI branch later
	RegisterReactiveGUIListeners(bus, g, commandChan, logArea, buttons, playerLabel, mainWindow)

	// Start command proccesing goroutine
	go func() {
		for cmd := range commandChan {
			g.Handle(cmd)
		}
	}()

	g.StartGame()
	mainWindow.ShowAndRun()
}
