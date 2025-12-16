package tile

import (
	"monopoly/helper"
	"monopoly/player"
)

type Street struct {
	PropertyTile
	housePrice            int
	houses                int
	priceIncreasePerHouse int
	color                 string
	hotelOwned            bool
}

func (street *Street) GetName() string {
	return street.Name
}

func (street *Street) GetPosition() int {
	return street.Position
}

func (street *Street) OnLand(player *player.Player) {

}

func (street *Street) GetPrice() int {
	return street.buyPrice
}

func (street *Street) GetRent(tiles []Tile) int {
	var rent = street.rent

	if street.GetHouseAmount() > 0 {
		return rent + street.priceIncreasePerHouse*street.GetHouseAmount()
	} else {
		return rent
	}
}

func NewStreetTile(buyPrice int, housePrice int, priceIncreasePerHouse int, color string, rent int, mortgageValue int) Property {
	return &Street{
		PropertyTile: PropertyTile{
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
	}
}

func (street *Street) SetHouseAmount(amount int) {
	street.houses = helper.Clamp(amount, 0, 4)
}

func (street *Street) GetHouseAmount() int {
	return street.houses
}

func (street *Street) buyHouse() {
	var curHouseAmount = street.GetHouseAmount()
	var newHouseAmount = curHouseAmount + 1

	street.SetHouseAmount(newHouseAmount)
}

func (street *Street) sellHouse() {
	var curHouseAmount = street.GetHouseAmount()
	var newHouseAmount = curHouseAmount - 1

	street.SetHouseAmount(newHouseAmount)
}

func (street *Street) buyHotel() {
	street.hotelOwned = true
	street.SetHouseAmount(0)
}

func (street *Street) sellHotel() {
	street.hotelOwned = false
	street.SetHouseAmount(4)
}
