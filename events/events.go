package events

import (
	"monopoly/common"
)

type PaidRentPayload struct {
	PlayerName  string
	Owner       string
	TileName    string
	Rent        int
	PlayerMoney int
}

type UpdateMoneyPayload struct {
	PlayerName string
	Money      int
}

type StartTurnPayload struct {
	PlayerName      string
	Money           int
	OwnedProperties []common.Tile
	Color           string
}

type JailedPayload struct {
	PlayerName  string
	JailedTurns int
}

type ForcedPayToExitJailPayload struct {
	PlayerName string
	Price      int
}

type PromptRollDicePayload struct {
	PlayerName string
}

type BoughtPropertyPayload struct {
	PlayerName string
	TileName   string
	Price      int
}

type DeclinedBuyPayload struct {
	PlayerName string
	TileName   string
	Amount     int
}

type LandedOnTilePayload struct {
	PlayerName string
	TileName   string
}

type LandedOnUnownedPropertyPayload struct {
	PlayerName string
	TileName   string
	Price      int
}

type LandedOnOwnPropertyPayload struct {
	PlayerName string
	TileName   string
}

type LandedOnGoToJailPayload struct {
	PlayerName string
	TileName   string
}

type LandedOnJailPayload struct {
	PlayerName string
	TileName   string
}

type LandedOnTaxPayload struct {
	PlayerName string
	TileName   string
	TaxAmount  int
}

type LandedOnFreeParkingPayload struct {
	PlayerName string
}

type RolledDicePayload struct {
	PlayerName string
	Dice       []int
}

type LandedOnGoPayload struct {
	PlayerName string
	TileName   string
}

type LandedOnChancePayload struct {
	PlayerName string
	TileName   string
}

type LandedOnCommunityChestPayload struct {
	PlayerName string
	TileName   string
}

type TaxPayload struct {
	PlayerName  string
	TileName    string
	TaxAmount   int
	PlayerMoney int
}

type InputPromptPayload struct {
	PlayerName string
	TileName   string
	Options    []any
}

type CantAffordPayload struct {
	Playername string
	TileName   string
	Price      int
}
