package tile

import (
	"monopoly/events"
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

func (goTile *JailTile) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	event := events.GameEvent{
		Type:       events.EventLandedOnJail,
		TileName:   goTile.GetName(),
		PlayerName: player.GetName(),
		Details:    "Landed on Jail tile",
	}

	return []events.GameEvent{event}
}

func NewJailTile(position int) Tile {
	return &JailTile{
		tile: BaseTile{
			Position: position,
			Name:     "Jail",
		},
	}
}
