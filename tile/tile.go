package tile

type BaseTile struct {
	Position int
	Name     string
}

func (b *BaseTile) GetPosition() int {
	return b.Position
}

func (b *BaseTile) GetName() string {
	return b.Name
}
