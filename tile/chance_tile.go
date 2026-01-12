package tile

import (
	"monopoly/card"
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

func (chanceTile *ChanceTile) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	event := events.GameEvent{
		Type: events.EventLandedOnChance,
		Payload: events.LandedOnChancePayload{
			PlayerName: player.GetName(),
			TileName:   chanceTile.GetName(),
		},
	}

	return []events.GameEvent{event}
}

func NewChanceTile(position int, name string) Tile {
	return &ChanceTile{
		BaseTile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}
