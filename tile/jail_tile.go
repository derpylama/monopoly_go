package tile

import (
	"monopoly/player"
)

type jailTile struct {
	tile    BaseTile
	players []player.Player
}

func tryToEscape(rolledDice []int) {

}

func (goTile *jailTile) GetName() string {
	return goTile.tile.Name
}

func (goTile *jailTile) GetPosition() int {
	return goTile.tile.Position
}

func (goTile *jailTile) OnLand(player *player.Player) {

}

func NewJailTile(position int) Tile {
	return &jailTile{
		tile: BaseTile{
			Position: position,
			Name:     "Jail",
		},
	}
}
