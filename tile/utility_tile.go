package tile

import (
	"monopoly/events"
	"monopoly/player"
	"strconv"
)

type Utility struct {
	PropertyTile
	baseMultiplier     int
	monopolyMultiplier int
}

func NewUtilityTile(buyPrice int, name string, mortgageValue int, baseMultiplier int, monopolyMultiplier int, position int) Tile {
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

func (utility *Utility) GetName() string {
	return utility.Name
}

func (utility *Utility) GetPosition() int {
	return utility.Position
}

func (utility *Utility) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	if utility.IsOwned() {
		if utility.GetOwner() != player {
			rent := utility.GetRent(tiles, dice)

			player.PayRent(utility.GetOwner(), rent)

			event := events.GameEvent{
				Type: events.EventPaidRent,
				Payload: events.PaidRentPayload{
					PlayerName: player.GetName(),
					TileName:   utility.GetName(),
					Details:    "Paid rent of " + strconv.Itoa(rent) + " to " + utility.GetOwner().GetName(),
					Amount:     rent,
				},
			}

			return []events.GameEvent{event}
		} else {
			event := events.GameEvent{
				Type: events.EventLandedOnOwnProperty,
				Payload: events.LandedOnOwnPropertyPayload{
					PlayerName: player.GetName(),
					TileName:   utility.GetName(),
				},
			}

			return []events.GameEvent{event}
		}
	} else {
		event := events.GameEvent{
			Type: events.EventLandedOnUnownedProperty,
			Payload: events.LandedOnUnownedPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   utility.GetName(),
				Amount:     utility.GetPrice(),
			},
		}

		return []events.GameEvent{event}

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

func (utility *Utility) GetRent(tiles []Tile, rolledDice []int) int {

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
