package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/helper"
	"monopoly/player"
)

type Color string

const (
	Brown     Color = "Brown"
	LightBlue Color = "Light Blue"
	Pink      Color = "Pink"
	Orange    Color = "Orange"
	Red       Color = "Red"
	Yellow    Color = "Yellow"
	Green     Color = "Green"
	DarkBlue  Color = "Dark Blue"
)

type Street struct {
	PropertyTile
	housePrice            int
	houses                int
	priceIncreasePerHouse []int
	color                 Color
	hotelOwned            bool
	hotelRent             int
}

func (street *Street) GetName() string {
	return street.Name
}

func (street *Street) GetPosition() int {
	return street.Position
}

func (street *Street) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *events.Bus) {

	bus.Publish(events.GameEvent{
		Type: events.LandedOnStreet,
		Payload: events.LandedOnTilePayload{
			PlayerName: player.GetName(),
			TileName:   street.GetName(),
		},
	})

	if !street.IsOwned() {
		bus.Publish(events.GameEvent{
			Type: events.LandedOnUnownedProperty,
			Payload: events.LandedOnUnownedPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   street.GetName(),
				Price:      street.GetPrice()},
		})
		return
	}

	if street.GetOwner() == player {
		return
	}

	rent := street.GetRent(tiles, dice)

	player.PayRent(street.GetOwner(), rent)

	bus.Publish(events.GameEvent{
		Type: events.PaidRent,
		Payload: events.PaidRentPayload{
			PlayerName: player.GetName(),
			TileName:   street.GetName(),
			Owner:      street.GetOwner().GetName(),
			Rent:       rent,
		},
	})
}

func (street *Street) GetPrice() int {
	return street.buyPrice
}

func (street *Street) GetRent(tiles []common.Tile, dice []int) int {
	var rent = street.rent

	if street.GetHouseAmount() > 0 {
		return street.priceIncreasePerHouse[street.GetHouseAmount()-1]
	} else {
		return rent
	}
}

func (street *Street) GetOwner() *player.Player {
	return street.ownedBy
}

func (street *Street) SetOwner(player *player.Player) {
	street.ownedBy = player
}

func (street *Street) IsOwned() bool {
	if street.ownedBy == nil {
		return false
	} else {
		return true
	}
}

func (street *Street) GetColor() Color {
	return street.color
}

func (street *Street) BuyProperty(player *player.Player, bus *events.Bus) {
	street.SetOwner(player)

	if player.CanAfford(street.GetPrice()) {
		player.Pay(street.GetPrice())

		bus.Publish(events.GameEvent{
			Type: events.BoughtProperty,
			Payload: events.BoughtPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   street.GetName(),
				Price:      street.GetPrice(),
			},
		})
	} else {
		bus.Publish(events.GameEvent{
			Type: events.CantAfford,
			Payload: events.CantAffordPayload{
				Playername: player.GetName(),
				TileName:   street.GetName(),
				Price:      street.GetPrice(),
			},
		})
	}

}

func NewStreetTile(buyPrice int, housePrice int, priceIncreasePerHouse []int, color Color, rent int, mortgageValue int, hotelRent int, name string, position int) Property {
	return &Street{
		PropertyTile: PropertyTile{
			BaseTile: BaseTile{
				Name:     name,
				Position: position,
			},
			buyPrice:      buyPrice,
			rent:          rent,
			ownedBy:       nil,
			mortgageValue: mortgageValue,
			isMortgage:    false,
		},
		housePrice:            housePrice,
		houses:                0,
		priceIncreasePerHouse: priceIncreasePerHouse,
		color:                 color,
		hotelOwned:            false,
		hotelRent:             hotelRent,
	}
}

func (street *Street) setHouseAmount(amount int) {
	street.houses = helper.Clamp(amount, 0, 4)
}

func (street *Street) GetHouseAmount() int {
	return street.houses
}

func (street *Street) GetHousePrice() int {
	return street.housePrice
}

func (street *Street) BuyHouse() {
	var curHouseAmount = street.GetHouseAmount()
	var newHouseAmount = curHouseAmount + 1

	street.setHouseAmount(newHouseAmount)
}

func (street *Street) SellHouse() {
	var curHouseAmount = street.GetHouseAmount()
	var newHouseAmount = curHouseAmount - 1

	street.setHouseAmount(newHouseAmount)
}

func (street *Street) BuyHotel() {
	street.hotelOwned = true
	street.setHouseAmount(0)
}

func (street *Street) SellHotel() {
	street.hotelOwned = false
	street.setHouseAmount(4)
}
