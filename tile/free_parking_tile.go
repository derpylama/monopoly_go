package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

type FreeParking struct {
	BaseTile
}

func NewFreeParkingTile(name string, position int) common.Tile {
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

func (freeParking *FreeParking) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *events.Bus) {
	bus.Publish(events.GameEvent{
		Type: events.LandedOnFreeParking,
		Payload: events.LandedOnFreeParkingPayload{
			PlayerName: player.GetName(),
		},
	})

}
