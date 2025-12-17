package tile

import (
	player "monopoly/player"
)

type PropertyTile struct {
	BaseTile
	buyPrice      int
	ownedBy       *player.Player
	rent          int
	mortgageValue int
	isMortgage    bool
}

type Property interface {
	Tile
	GetPrice() int
	GetRent(tiles []Tile, rolledDice []int) int
}

func (tile *PropertyTile) Mortgage() int {
	tile.isMortgage = true
	return tile.mortgageValue
}
