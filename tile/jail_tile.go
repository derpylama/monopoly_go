package tile

import (
	"monopoly/common"
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

func (goTile *JailTile) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *common.Bus) {
	bus.Publish(common.GameEvent{
		Type: common.LandedOnJail,
		Payload: events.LandedOnJailPayload{
			PlayerName: player.GetName(),
			TileName:   goTile.GetName(),
		},
	})
}

func NewJailTile(position int) common.Tile {
	return &JailTile{
		tile: BaseTile{
			Position: position,
			Name:     "Jail",
		},
	}
}
