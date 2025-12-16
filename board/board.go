package board

import (
	"monopoly/tile"
)

type Board struct {
	tiles []tile.Tile
}

func (board *Board) initBoard() {

	board.tiles = []tile.Tile{
		tile.NewGoTile(0, 200),
		tile.NewJailTile(10),
	}
}

func NewBoard() *Board {
	b := &Board{}
	b.initBoard()
	return b
}

// Exported accessor
func (b *Board) Tiles() []tile.Tile {
	return b.tiles
}
