package tile

import (
	"monopoly/card"
	"monopoly/player"
)

type CommunityChest struct {
	tile  BaseTile
	cards []card.Card
}

func (commnityChest *CommunityChest) GetName() string {
	return commnityChest.tile.Name
}

func (commnityChest *CommunityChest) GetPosition() int {
	return commnityChest.tile.Position
}

func (commnityChest *CommunityChest) OnLand(player *player.Player) {

}

func NewCommunityChest(position int, name string) Tile {
	return &CommunityChest{
		tile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}
