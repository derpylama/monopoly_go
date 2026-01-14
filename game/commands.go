package game

type CommandType string

const (
	CmdRollDice    CommandType = "roll_dice"
	CmdEndTurn     CommandType = "end_turn"
	CmdBuyProperty CommandType = "buy_property"
)

type GameCommand struct {
	Type       CommandType
	PlayerName string
	TileName   string
}
