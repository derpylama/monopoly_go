package game

import (
	"monopoly/player"
)

type Game struct {
	players       []player.Player
	currentPlayer int
}
