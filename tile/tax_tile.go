package tile

type TaxTile struct {
	tile      Tile
	TaxAmount int
}

func (taxTile *TaxTile) getTax() int {
	return taxTile.TaxAmount
}

// func NewTaxTile(position int, taxAmount int, name string) *TaxTile {
// 	return &TaxTile{
// 		tile: Tile{
// 			Position: position,
// 			Name:     name,
// 		},
// 		TaxAmount: taxAmount,
// 	}
// }
