package tile

import "monopoly/player"

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

func (freeParking *FreeParking) OnLand(*player.Player) {

}
