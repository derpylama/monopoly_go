package logger

// import (
// 	"fmt"
// 	"monopoly/events"
// 	"strconv"
// )

// func LogOnLandInfo(event events.GameEvent) {
// 	ClearScreen()
// 	fmt.Println("Player " + event.Payload.PlayerName + " landed on " + event.TileName + ". Details: " + event.Details + "\n")
// }

// func LogEvent(event []events.GameEvent) {
// 	ClearScreen()

// 	for _, e := range event {
// 		fmt.Println("Event: " + string(e.Type) + "\nPlayer: " + e.PlayerName + "\nTile: " + e.TileName + "\nDetails: " + e.Details)
// 	}
// }

// func LogOwnedProperties(playerName string, properties []string) {
// 	ClearScreen()

// 	var count int = 0

// 	fmt.Println("Player " + playerName + " owns the following properties:")
// 	for _, property := range properties {
// 		fmt.Println(strconv.Itoa(count) + " - " + property)
// 		count++
// 	}
// 	fmt.Println()
// }

// func ClearScreen() {
// 	fmt.Print("\033[H\033[2J")
// }

import (
	"fmt"
)

type Logger struct {
	enabled bool
}

func New(enabled bool) *Logger {
	return &Logger{enabled: enabled}
}

func (l *Logger) log(msg string) {
	if !l.enabled {
		return
	}
	fmt.Println(msg)
}
