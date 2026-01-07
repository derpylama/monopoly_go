package tile

import (
	"monopoly/player"
)

type JailTile struct {
	tile    BaseTile
	players []player.Player
}

func tryToEscape(rolledDice []int) {

}

func (goTile *JailTile) GetName() string {
	return goTile.tile.Name
}

func (goTile *JailTile) GetPosition() int {
	return goTile.tile.Position
}

func (goTile *JailTile) OnLand(player *player.Player) {

}

func NewJailTile(position int) Tile {
	return &JailTile{
		tile: BaseTile{
			Position: position,
			Name:     "Jail",
		},
	}
}
