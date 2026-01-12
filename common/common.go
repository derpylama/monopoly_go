package common

import (
	"monopoly/events"
	"monopoly/player"
)

type Tile interface {
	GetPosition() int
	GetName() string
	OnLand(player *player.Player, tiles []Tile, dice []int, bus *events.Bus)
}
