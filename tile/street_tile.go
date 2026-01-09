package tile

import (
	"fmt"
	"monopoly/events"
	"monopoly/helper"
	"monopoly/inputhandler"

	"monopoly/player"
)

type Street struct {
	PropertyTile
	housePrice            int
	houses                int
	priceIncreasePerHouse []int
	color                 string
	hotelOwned            bool
	hotelRent             int
}

func (street *Street) GetName() string {
	return street.Name
}

func (street *Street) GetPosition() int {
	return street.Position
}

func (street *Street) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {

	eventList := []events.GameEvent{}

	if street.IsOwned() && street.GetOwner() != player {
		// if the property is owned then get rent and pay
		rent := street.GetRent(tiles, dice)

		player.PayRent(street.GetOwner(), rent)

		eventList = append(eventList, events.GameEvent{
			Type:       events.EventPaidRent,
			PlayerName: player.GetName(),
			Owner:      street.GetOwner().GetName(),
			TileName:   street.GetName(),
			Amount:     rent,
		})

	}

	if !street.IsOwned() && player.CanAfford(street.GetPrice()) {
		if inputhandler.PlayerWantsToBuyProperty(player.GetName(), street.GetName(), street.GetPrice()) {
			street.SetOwner(player)
			player.Pay(street.GetPrice())

			fmt.Println(player.GetName(), "bought", street.GetName(), "for", street.GetPrice())

			eventList = append(eventList, events.GameEvent{
				Type:       events.EventLandedOnUnownedProperty,
				PlayerName: player.GetName(),
				TileName:   street.GetName(),
				Amount:     street.GetPrice(),
			})

			eventList = append(eventList, events.GameEvent{
				Type:       events.EventBoughtProperty,
				PlayerName: player.GetName(),
				TileName:   street.GetName(),
				Amount:     street.GetPrice(),
			})
		} else {
			eventList = append(eventList, events.GameEvent{
				Type:       events.EventDeclinedBuy,
				PlayerName: player.GetName(),
				TileName:   street.GetName(),
			})

		}
	}
	return eventList
}

func (street *Street) GetPrice() int {
	return street.buyPrice
}

func (street *Street) GetRent(tiles []Tile, dice []int) int {
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

func NewStreetTile(buyPrice int, housePrice int, priceIncreasePerHouse []int, color string, rent int, mortgageValue int, hotelRent int, name string, position int) Tile {
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
