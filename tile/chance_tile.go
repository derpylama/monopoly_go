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

func (chanceTile *ChanceTile) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *common.Bus) {

	bus.Publish(common.GameEvent{
		Type: common.LandedOnChance,
		Payload: events.LandedOnChancePayload{
			PlayerName: player.GetName(),
			TileName:   chanceTile.GetName(),
		},
	})
}

func (chance *ChanceTile) BuyProperty(player *player.Player, bus *common.Bus) {

}

func NewChanceTile(position int, name string) common.Tile {
	return &ChanceTile{
		BaseTile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}
