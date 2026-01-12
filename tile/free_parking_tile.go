package tile

import (
	"monopoly/events"
	"monopoly/player"
)

type FreeParking struct {
	BaseTile
}

func NewFreeParkingTile(name string, position int) Tile {
	return &FreeParking{
		BaseTile: BaseTile{
			Name:     name,
			Position: position,
		},
	}
}

func (freeParking *FreeParking) GetName() string {
	return freeParking.Name
}

func (freeParking *FreeParking) GetPosition() int {
	return freeParking.Position
}

func (freeParking *FreeParking) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	return []events.GameEvent{
		{
			Type: events.EventLandedOnFreeParking,
			Payload: events.LandedOnFreeParkingPayload{
				PlayerName: player.GetName(),
			},
		},
	}
}
