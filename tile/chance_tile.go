package tile

import (
	"monopoly/card"
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

type ChanceTile struct {
	BaseTile
	cards []card.Card
}

func (chanceTile *ChanceTile) GetName() string {
	return chanceTile.Name
}

func (chanceTile *ChanceTile) GetPosition() int {
	return chanceTile.Position
}

func (chanceTile *ChanceTile) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *events.Bus) {

	bus.Publish(events.GameEvent{
		Type: events.LandedOnChance,
		Payload: events.LandedOnChancePayload{
			PlayerName: player.GetName(),
			TileName:   chanceTile.GetName(),
		},
	})
}

func NewChanceTile(position int, name string) common.Tile {
	return &ChanceTile{
		BaseTile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}
