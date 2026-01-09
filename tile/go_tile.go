package tile

import (
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

func (goTile *goTile) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	event := events.GameEvent{
		PlayerName: player.GetName(),
		TileName:   goTile.GetName(),
		Type:       events.EventLandedOnGo,
		Details:    "Landed on Go tile",
	}

	events := []events.GameEvent{event}

	return events

}

func NewGoTile(position int, goMoneyAmount int) Tile {
	return &goTile{
		tile: BaseTile{
			Position: position,
			Name:     "Go"},
		goMoneyAmount: goMoneyAmount,
	}
}
