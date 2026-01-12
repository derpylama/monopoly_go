package tile

import (
	"monopoly/card"
	"monopoly/common"
	"monopoly/events"
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

func (communityChest *CommunityChest) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *events.Bus) {
	bus.Publish(events.GameEvent{
		Type: events.LandedOnTile,
		Payload: events.LandedOnTilePayload{
			PlayerName: player.GetName(),
			TileName:   communityChest.GetName(),
		},
	})
}

func NewCommunityChestTile(position int, name string) common.Tile {
	return &CommunityChest{
		tile: BaseTile{
			Position: position,
			Name:     name,
		},
	}
}

func (communityChest *CommunityChest) initCards() {

}
