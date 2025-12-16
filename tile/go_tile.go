package tile

import "monopoly/player"

type goTile struct {
	tile          BaseTile
	goMoneyAmount int
}

func (goTile *goTile) passGo() int {
	return goTile.goMoneyAmount
}

func (goTile *goTile) GetName() string {
	return goTile.tile.Name
}

func (goTile *goTile) GetPosition() int {
	return goTile.tile.Position
}

func (goTile *goTile) OnLand(player *player.Player) {

}

func NewGoTile(position int, goMoneyAmount int) Tile {
	return &goTile{
		tile: BaseTile{
			Position: position,
			Name:     "Go"},
		goMoneyAmount: goMoneyAmount,
	}
}
