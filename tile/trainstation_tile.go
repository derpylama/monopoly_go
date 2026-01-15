package tile

import (
	"monopoly/common"
	"monopoly/events"
	"monopoly/player"
)

type TrainStation struct {
	PropertyTile
	rentPerStation int
}

func NewTrainStation(buyPrice int, mortgageValue int, name string, position int, rentPerStation int) Property {
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

func (trainStation *TrainStation) BuyProperty(player *player.Player, bus *common.Bus) {
	if player.CanAfford(trainStation.GetPrice()) {
		trainStation.SetOwner(player)
		player.Pay(trainStation.GetPrice())

		bus.Publish(common.GameEvent{
			Type: common.BoughtProperty,
			Payload: events.BoughtPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   trainStation.GetName(),
				Price:      trainStation.GetPrice(),
			},
		})

		bus.Publish(common.GameEvent{
			Type: common.UpdateMoney,
			Payload: events.UpdateMoneyPayload{
				PlayerName: player.GetName(),
				Money:      player.GetMoney(),
			},
		})
	} else {
		bus.Publish(common.GameEvent{
			Type: common.CantAfford,
			Payload: events.CantAffordPayload{
				Playername: player.GetName(),
				TileName:   trainStation.GetName(),
				Price:      trainStation.GetPrice(),
			},
		})
	}

}

func (trainStation *TrainStation) GetName() string {
	return trainStation.Name
}

func (trainStation *TrainStation) GetPosition() int {
	return trainStation.Position
}

func (trainStation *TrainStation) OnLand(player *player.Player, tiles []common.Tile, dice []int, bus *common.Bus) {
	if trainStation.IsOwned() {
		if trainStation.GetOwner() != player {
			rent := trainStation.GetRent(tiles, dice)
			bus.Publish(common.GameEvent{
				Type: common.PaidRent,
				Payload: events.PaidRentPayload{
					PlayerName: player.GetName(),
					Owner:      trainStation.GetOwner().GetName(),
					TileName:   trainStation.GetName(),
					Rent:       rent,
				},
			})
		}
	} else {
		bus.Publish(common.GameEvent{
			Type: common.LandedOnUnownedProperty,
			Payload: events.LandedOnUnownedPropertyPayload{
				PlayerName: player.GetName(),
				TileName:   trainStation.GetName(),
				Price:      trainStation.GetPrice(),
			},
		})

	}
}

func (trainStation *TrainStation) GetPrice() int {
	return trainStation.buyPrice
}

func (trainStation *TrainStation) GetRent(tiles []common.Tile, rolledDice []int) int {

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
