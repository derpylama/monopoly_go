package tile

import (
	"monopoly/card"
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

func (chanceTile *ChanceTile) OnLand(player *player.Player) {

}

func NewChanceTile(position int, name string) Tile {
	return &ChanceTile{
		BaseTile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}
