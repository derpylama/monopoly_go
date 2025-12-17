package tile

import (
	"monopoly/helper"
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

func (street *Street) OnLand(player *player.Player) {

}

func (street *Street) GetPrice() int {
	return street.buyPrice
}

func (street *Street) GetRent(tiles []Tile, rolledDice []int) int {
	var rent = street.rent

	if street.GetHouseAmount() > 0 {
		return street.priceIncreasePerHouse[street.GetHouseAmount()-1]
	} else {
		return rent
	}
}

func NewStreetTile(buyPrice int, housePrice int, priceIncreasePerHouse []int, color string, rent int, mortgageValue int, hotelRent int, name string, position int) Property {
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

func (street *Street) buyHouse() {
	var curHouseAmount = street.GetHouseAmount()
	var newHouseAmount = curHouseAmount + 1

	street.setHouseAmount(newHouseAmount)
}

func (street *Street) sellHouse() {
	var curHouseAmount = street.GetHouseAmount()
	var newHouseAmount = curHouseAmount - 1

	street.setHouseAmount(newHouseAmount)
}

func (street *Street) buyHotel() {
	street.hotelOwned = true
	street.setHouseAmount(0)
}

func (street *Street) sellHotel() {
	street.hotelOwned = false
	street.setHouseAmount(4)
}
