package tile

import (
	"monopoly/events"
	"monopoly/player"
)

type Tile interface {
	GetPosition() int
	GetName() string
	OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent
}

type BaseTile struct {
	Position int
	Name     string
}

func (b *BaseTile) GetPosition() int {
	return b.Position
}

func (b *BaseTile) GetName() string {
	return b.Name
}
