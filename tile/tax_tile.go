package tile

import (
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

func (taxTile *TaxTile) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	playerMoney := player.GetMoney()
	playerMoney -= taxTile.getTax()

	player.SetMoney(playerMoney)

	event := events.GameEvent{
		Type: events.EventPaidTax,
		Payload: events.LandedOnTaxPayload{
			PlayerName: player.GetName(),
			TileName:   taxTile.GetName(),
			Amount:     taxTile.getTax(),
		},
	}

	return []events.GameEvent{event}
}

func (taxTile *TaxTile) GetPosition() int {
	return taxTile.Position
}

func (taxTile *TaxTile) GetTaxAmount() int {
	return taxTile.taxAmount
}
