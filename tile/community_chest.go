package tile

import (
	"monopoly/card"
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

func (communityChest *CommunityChest) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	event := events.GameEvent{
		PlayerName: player.GetName(),
		TileName:   communityChest.GetName(),
		Type:       events.EventLandedOnCommunityChest,
		Details:    "Landed on community chest tile",
	}

	return []events.GameEvent{event}
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
