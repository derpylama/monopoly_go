package tile

import (
	"fmt"
	"monopoly/common"
	"monopoly/player"
)

type GoToJail struct {
	BaseTile
}

func NewGoToJailTile(name string, position int) common.Tile {
	return &GoToJail{
		BaseTile: BaseTile{
			Name:     name,
			Position: position,
		},
	}
}

func (jail *GoToJail) GetName() string {
	return jail.Name
}

func (jail *GoToJail) GetPosition() int {
	return jail.Position
}

func (jail *GoToJail) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *common.Bus) {
	if !player.GetJailStatus() {
		jailPos, ok := common.GetTilePosByName("Jail", tiles)

		if ok {
			player.Teleport(jailPos)
			player.ToggleIsJailed()
			player.IncrementJailedTurns()

		} else {
			fmt.Println("Jail Tile not found")
		}
	}
}
