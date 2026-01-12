package game

type CommandType string

const (
	CmdRollDice CommandType = "roll_dice"
)

type GameCommand struct {
	Type       CommandType
	PlayerName string
	TileName   string
}
