package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

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

func (goTile *goTile) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *common.Bus) {
	bus.Publish(common.GameEvent{
		Type: common.LandedOnGo,
		Payload: events.LandedOnGoPayload{
			PlayerName: player.GetName(),
			TileName:   goTile.GetName(),
		},
	})

}

func NewGoTile(position int, goMoneyAmount int) common.Tile {
	return &goTile{
		tile: BaseTile{
			Position: position,
			Name:     "Go"},
		goMoneyAmount: goMoneyAmount,
	}
}
