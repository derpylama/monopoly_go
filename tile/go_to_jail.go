package tile

import (
	"monopoly/events"
	"monopoly/player"
)

type GoToJail struct {
	BaseTile
}

func (jail *GoToJail) setPlayerToJail(player *player.Player) {

}

func NewGoToJailTile(name string, position int) Tile {
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

func (jail *GoToJail) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	event := events.GameEvent{
		PlayerName: player.GetName(),
		Type:       events.EventLandedOnGoToJail,
		Details:    "Landed on Go To Jail",
		TileName:   jail.GetName(),
	}
	return []events.GameEvent{event}
}
