package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

type GoToJail struct {
	BaseTile
}

func (jail *GoToJail) setPlayerToJail(player *player.Player) {

}

func NewGoToJailTile(name string, position int) common.Tile {
	return &GoToJail{
		BaseTile: BaseTile{
			Name:     name,
			Position: position,
		},
	}
}

func (jail *GoToJail) GetName() string {
	return jail.Name
}

func (jail *GoToJail) GetPosition() int {
	return jail.Position
}

func (jail *GoToJail) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *events.Bus) {

}
