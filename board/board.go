package board

import (
	"monopoly/common"
	"monopoly/helper"
	"monopoly/tile"
)

type Board struct {
	tiles []common.Tile
}

func (board *Board) initBoard() {

	board.tiles = []common.Tile{
		tile.NewGoTile(0, 200),

		// tile.NewStreetTile(buyPrice int, price per house []int, color string, rent int, mortageValue int, hotelRent int, name string, position int)

		//Brown streets
		tile.NewStreetTile(60, 50, []int{10, 30, 90, 160}, tile.Brown, 2, 30, 250, "Mediterranean Avenue", 1),
		tile.NewCommunityChestTile(2, "Community Chest"),
		tile.NewStreetTile(60, 50, []int{20, 60, 180, 320}, tile.Brown, 4, 30, 450, "Baltic Avenue", 3),

		tile.NewTaxTile(4, 200, "Income Tax"),

		tile.NewTrainStation(200, 100, "Reading Railroad", 5, 50),

		//Light blue steets
		tile.NewStreetTile(100, 50, []int{30, 90, 270, 400}, tile.LightBlue, 6, 50, 550, "Oriental Avenue", 6),

		tile.NewChanceTile(7, "Chance"),

		tile.NewStreetTile(100, 50, []int{30, 90, 270, 400}, tile.LightBlue, 6, 50, 550, "Vermont Avenue", 8),
		tile.NewStreetTile(120, 50, []int{40, 100, 300, 450}, tile.LightBlue, 8, 60, 600, "Conneticut Avenue", 9),

		tile.NewJailTile(10),

		//Pink streets
		tile.NewStreetTile(140, 100, []int{50, 150, 450, 625}, tile.Pink, 10, 70, 750, "St. Charles Place", 11),

		tile.NewUtilityTile(150, "Electric Company", 75, 4, 10, 12),

		tile.NewStreetTile(140, 100, []int{50, 150, 450, 625}, tile.Pink, 10, 70, 750, "States Avenue", 13),
		tile.NewStreetTile(160, 100, []int{60, 180, 500, 700}, tile.Pink, 12, 80, 900, "Virgina Avenue", 14),

		tile.NewTrainStation(200, 100, "Pennsylvania Railroad", 15, 50),

		//Orange streets
		tile.NewStreetTile(180, 100, []int{70, 200, 550, 750}, tile.Orange, 14, 90, 950, "St. James Place", 16),

		tile.NewCommunityChestTile(17, "Community Chest"),

		tile.NewStreetTile(180, 100, []int{70, 200, 550, 750}, tile.Orange, 14, 90, 950, "Tennessee Avenue", 18),
		tile.NewStreetTile(200, 100, []int{80, 220, 600, 800}, tile.Orange, 16, 100, 1000, "New York Avenue", 19),

		tile.NewFreeParkingTile("Free Parking", 20),

		//Red streets
		tile.NewStreetTile(220, 150, []int{90, 250, 700, 875}, tile.Red, 18, 110, 1050, "Kentucky Avenue", 21),

		tile.NewChanceTile(22, "Chance"),

		tile.NewStreetTile(220, 150, []int{90, 250, 700, 875}, tile.Red, 18, 110, 1050, "Indiana Avenue", 23),
		tile.NewStreetTile(240, 150, []int{100, 300, 750, 925}, tile.Red, 20, 120, 1100, "Illinois Avenue", 24),

		tile.NewTrainStation(200, 100, "B&O Railroad", 25, 50),

		//Yellow streets
		tile.NewStreetTile(260, 150, []int{110, 330, 800, 975}, tile.Yellow, 22, 130, 1150, "Atlantic Avenue", 26),
		tile.NewStreetTile(260, 150, []int{110, 330, 800, 975}, tile.Yellow, 22, 130, 1150, "Ventnor Avenue", 27),

		tile.NewUtilityTile(150, "Water Works", 75, 4, 10, 28),

		tile.NewStreetTile(280, 150, []int{120, 360, 850, 1025}, tile.Yellow, 24, 140, 1200, "Marvin Gardens", 29),

		tile.NewGoToJailTile("Go To Jail", 30),

		//Green streets
		tile.NewStreetTile(300, 200, []int{130, 390, 900, 1100}, tile.Green, 26, 150, 1275, "Pacific Avenue", 31),
		tile.NewStreetTile(300, 300, []int{130, 390, 900, 1100}, tile.Green, 26, 150, 1275, "North Carolina Avenue", 32),

		tile.NewCommunityChestTile(33, "Community Chest"),

		tile.NewStreetTile(320, 200, []int{150, 450, 1000, 1200}, tile.Green, 28, 160, 1400, "Pennsylvania", 34),

		tile.NewTrainStation(200, 100, "Short Line", 35, 50),

		tile.NewChanceTile(36, "Chance"),

		//Blue streets
		tile.NewStreetTile(350, 200, []int{175, 500, 1100, 1300}, tile.DarkBlue, 35, 175, 1500, "Park Place", 37),

		tile.NewTaxTile(38, 100, "Luxury Tax"),

		tile.NewStreetTile(400, 200, []int{200, 600, 1400, 1700}, tile.DarkBlue, 50, 200, 2000, "Boardwalk", 39),
	}
}

func NewBoard() *Board {
	b := &Board{}
	b.initBoard()
	return b
}

// Exported accessor
func (b *Board) Tiles() []common.Tile {
	return b.tiles
}

func (board *Board) GetTile(position int) common.Tile {

	clampedVal := helper.Clamp(position, 0, 39)

	return board.tiles[clampedVal]
}

func (board *Board) StreetByColor(color tile.Color) []common.Tile {
	var streets []common.Tile

	for _, t := range board.tiles {
		if street, ok := t.(*tile.Street); ok {
			if street.GetColor() == color {
				streets = append(streets, street)
			}
		}
	}

	return streets
}

func (board *Board) GetTileByName(name string) (common.Tile, bool) {
	for i, t := range board.Tiles() {
		if t.GetName() == name {
			return board.GetTile(i), true
		}

	}
	return nil, false
}
