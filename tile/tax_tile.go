package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

type TaxTile struct {
	BaseTile
	taxAmount int
}

func (taxTile *TaxTile) getTax() int {
	return taxTile.taxAmount
}

func NewTaxTile(position int, taxAmount int, name string) common.Tile {
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

func (taxTile *TaxTile) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *events.Bus) {
	playerMoney := player.GetMoney()
	playerMoney -= taxTile.getTax()

	player.SetMoney(playerMoney)

	bus.Publish(events.GameEvent{
		Type: events.LandedOnTax,
		Payload: events.TaxPayload{
			PlayerName: player.GetName(),
			TileName:   taxTile.GetName(),
			TaxAmount:  taxTile.getTax(),
		},
	})

	player.Pay(taxTile.getTax())

	bus.Publish(events.GameEvent{
		Type: events.PaidTax,
		Payload: events.TaxPayload{
			PlayerName:  player.GetName(),
			TileName:    taxTile.GetName(),
			TaxAmount:   taxTile.getTax(),
			PlayerMoney: player.GetMoney(),
		},
	})
}

func (taxTile *TaxTile) GetPosition() int {
	return taxTile.Position
}

func (taxTile *TaxTile) GetTaxAmount() int {
	return taxTile.taxAmount
}
