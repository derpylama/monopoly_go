package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

type Utility struct {
	PropertyTile
	baseMultiplier     int
	monopolyMultiplier int
}

func NewUtilityTile(buyPrice int, name string, mortgageValue int, baseMultiplier int, monopolyMultiplier int, position int) Property {
	return &Utility{
		PropertyTile: PropertyTile{
			buyPrice:      buyPrice,
			mortgageValue: mortgageValue,
			BaseTile: BaseTile{
				Name:     name,
				Position: position,
			},
		},
		baseMultiplier:     baseMultiplier,
		monopolyMultiplier: monopolyMultiplier,
	}
}

func (utility *Utility) BuyProperty(player *player.Player, bus *common.Bus) {
	if player.CanAfford(utility.GetPrice()) {
		utility.SetOwner(player)
		player.Pay(utility.GetPrice())

		bus.Publish(common.GameEvent{
			Type: common.BoughtProperty,
			Payload: events.BoughtPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   utility.GetName(),
				Price:      utility.GetPrice(),
			},
		})

		bus.Publish(common.GameEvent{
			Type: common.UpdateMoney,
			Payload: events.UpdateMoneyPayload{
				PlayerName: player.GetName(),
				Money:      player.GetMoney(),
			},
		})
	} else {
		bus.Publish(common.GameEvent{
			Type: common.CantAfford,
			Payload: events.CantAffordPayload{
				Playername: player.GetName(),
				TileName:   utility.GetName(),
				Price:      utility.GetPrice(),
			},
		})
	}
}

func (utility *Utility) GetName() string {
	return utility.Name
}

func (utility *Utility) GetPosition() int {
	return utility.Position
}

func (utility *Utility) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *common.Bus) {
	if utility.IsOwned() {
		if utility.GetOwner() != player {
			rent := utility.GetRent(tiles, dice)
			bus.Publish(common.GameEvent{
				Type: common.PaidRent,
				Payload: events.PaidRentPayload{
					PlayerName: player.GetName(),
					Owner:      utility.GetOwner().GetName(),
					TileName:   utility.GetName(),
					Rent:       rent,
				},
			})
		}
	} else {
		bus.Publish(common.GameEvent{
			Type: common.LandedOnUnownedProperty,
			Payload: events.LandedOnUnownedPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   utility.GetName(),
				Price:      utility.GetPrice(),
			},
		})

	}

}

func (utility *Utility) GetPrice() int {
	return utility.buyPrice
}

func (utility *Utility) IsOwned() bool {
	if utility.ownedBy == nil {
		return false
	} else {
		return true
	}
}

func (utility *Utility) GetRent(tiles []common.Tile, rolledDice []int) int {

	var diceTotal int

	for _, dice := range rolledDice {
		diceTotal += dice
	}

	count := 0

	for _, tile := range tiles {
		if u, ok := tile.(*Utility); ok {
			if u.ownedBy == utility.ownedBy && utility != nil {
				count++
			}
		}
	}

	if count == 2 {
		return diceTotal * utility.monopolyMultiplier
	} else {
		return diceTotal * utility.baseMultiplier
	}
}

func (utility *Utility) GetOwner() *player.Player {
	return utility.ownedBy
}

func (utility *Utility) SetOwner(player *player.Player) {
	utility.ownedBy = player
}
