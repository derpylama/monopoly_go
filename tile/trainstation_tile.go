package tile

import (
	"monopoly/events"
	"monopoly/player"
)

type TrainStation struct {
	PropertyTile
	rentPerStation int
}

func NewTrainStation(buyPrice int, mortgageValue int, name string, position int, rentPerStation int) Tile {
	return &TrainStation{
		PropertyTile: PropertyTile{
			BaseTile: BaseTile{
				Name:     name,
				Position: position,
			},
			buyPrice:      buyPrice,
			mortgageValue: mortgageValue,
		},
		rentPerStation: rentPerStation,
	}
}

func (trainStation *TrainStation) IsOwned() bool {
	if trainStation.ownedBy == nil {
		return false
	} else {
		return true
	}
}

func (trainStation *TrainStation) GetName() string {
	return trainStation.Name
}

func (trainStation *TrainStation) GetPosition() int {
	return trainStation.Position
}

func (trainStation *TrainStation) OnLand(player *player.Player, tiles []Tile, dice []int) []events.GameEvent {
	eventList := []events.GameEvent{}

	if !trainStation.IsOwned() {
		eventList = append(eventList, events.GameEvent{
			Type:       events.EventLandedOnUnownedProperty,
			PlayerName: player.GetName(),
			TileName:   trainStation.GetName(),
			Details:    player.GetName() + " landed on unowned property " + trainStation.Name,
		})
	} else if trainStation.GetOwner() != player {
		rent := trainStation.GetRent(tiles, dice)
		player.PayRent(trainStation.GetOwner(), rent)

		eventList = append(eventList, events.GameEvent{
			Type:       events.EventPaidRent,
			PlayerName: player.GetName(),
			TileName:   trainStation.GetName(),
			Amount:     rent,
			Details:    player.GetName() + " paid rent of " + string(rent) + " to " + trainStation.GetOwner().GetName() + " for landing on " + trainStation.GetName(),
		})
	} else {
		eventList = append(eventList, events.GameEvent{
			Type:       events.EventPaidRent,
			PlayerName: player.GetName(),
			TileName:   trainStation.GetName(),
			Details:    player.GetName() + " landed on their own property " + trainStation.GetName(),
		})
	}

	return eventList

}

func (trainStation *TrainStation) GetPrice() int {
	return trainStation.buyPrice
}

func (trainStation *TrainStation) GetRent(tiles []Tile, rolledDice []int) int {

	count := 0

	for _, tile := range tiles {
		if u, ok := tile.(*TrainStation); ok {
			if u.ownedBy == trainStation.ownedBy && trainStation != nil {
				count++
			}
		}
	}

	return trainStation.rentPerStation * count
}

func (trainStation *TrainStation) GetOwner() *player.Player {
	return trainStation.ownedBy
}

func (trainStation *TrainStation) SetOwner(player *player.Player) {
	trainStation.ownedBy = player
}
