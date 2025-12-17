package tile

import "monopoly/player"

type TaxTile struct {
	BaseTile
	TaxAmount int
}

func (taxTile *TaxTile) getTax() int {
	return taxTile.TaxAmount
}

func NewTaxTile(position int, taxAmount int, name string) Tile {
	return &TaxTile{
		BaseTile: BaseTile{
			Position: position,
			Name:     name,
		},
		TaxAmount: taxAmount,
	}
}

func (taxTile *TaxTile) GetName() string {
	return taxTile.Name
}

func (taxTile *TaxTile) OnLand(player *player.Player) {

}

func (taxTile *TaxTile) GetPosition() int {
	return taxTile.Position
}
