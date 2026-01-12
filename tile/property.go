package tile

import (
	"monopoly/common"
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
	common.Tile
	GetPrice() int
	GetRent(tiles []common.Tile, rolledDice []int) int
	GetOwner() *player.Player
	SetOwner(*player.Player)
	IsOwned() bool
	BuyProperty(player *player.Player)
}

func (tile *PropertyTile) Mortgage() int {
	tile.isMortgage = true
	return tile.mortgageValue
}
