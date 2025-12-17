package tile

import "monopoly/player"

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

func (trainStation *TrainStation) GetName() string {
	return trainStation.Name
}

func (trainStation *TrainStation) GetPosition() int {
	return trainStation.Position
}

func (trainStation *TrainStation) OnLand(player *player.Player) {

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
