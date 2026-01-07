package tile

import (
	"monopoly/card"
	"monopoly/player"
)

type CommunityChest struct {
	tile  BaseTile
	cards []card.Card
}

func (communityChest *CommunityChest) GetName() string {
	return communityChest.tile.Name
}

func (communityChest *CommunityChest) GetPosition() int {
	return communityChest.tile.Position
}

func (commnityChest *CommunityChest) OnLand(player *player.Player) {

}

func NewCommunityChestTile(position int, name string) Tile {
	return &CommunityChest{
		tile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}

func (commityChest *CommunityChest) initCards() {

}
