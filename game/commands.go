package game

type CommandType string

const (
	CmdRollDice      CommandType = "roll_dice"
	CmdEndTurn       CommandType = "end_turn"
	CmdBuyProperty   CommandType = "buy_property"
	CmdPayToExitJail CommandType = "pay_to_exit_jail"
	CmdMortgage      CommandType = "mortgage"
	CmdUnMortgage    CommandType = "un_mortgage"
	CmdBuildHouse    CommandType = "build_house"
	CmdBuildHotel    CommandType = "build_hotel"
)

type GameCommand struct {
	Type       CommandType
	PlayerName string
	TileName   string
}
