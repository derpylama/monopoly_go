package tile

import "monopoly/player"

type TaxTile struct {
	BaseTile
	taxAmount int
}

func (taxTile *TaxTile) getTax() int {
	return taxTile.taxAmount
}

func NewTaxTile(position int, taxAmount int, name string) Tile {
	return &TaxTile{
		BaseTile: BaseTile{
			Position: position,
			Name:     name,
		},
		taxAmount: taxAmount,
	}
}

func (taxTile *TaxTile) GetName() string {
	return taxTile.Name
}

func (taxTile *TaxTile) OnLand(player *player.Player) {
	playerMoney := player.GetMoney()
	playerMoney -= taxTile.getTax()

	player.SetMoney(playerMoney)

}

func (taxTile *TaxTile) GetPosition() int {
	return taxTile.Position
}

func (taxTile *TaxTile) GetTaxAmount() int {
	return taxTile.taxAmount
}
