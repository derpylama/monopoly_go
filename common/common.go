package common

import (
	"monopoly/player"
	"sync"
)

type Tile interface {
	GetPosition() int
	GetName() string
	OnLand(player *player.Player, tiles []Tile, dice []int, bus *Bus)
}

// Bus is the central event dispatcher
type Bus struct {
	mu        sync.RWMutex
	listeners map[EventType][]Listener
}

// Type represents an event type (paid_rent, rolled_dice, etc.)
type Type string

// Event is the envelope passed through the bus
type Event struct {
	Type    Type
	Payload any
}

// Listener is a function that reacts to an event
type Listener func(GameEvent)

type EventType string

const (
	StartTurn               EventType = "start_turn"
	EndTurn                 EventType = "end_turn"
	PaidRent                EventType = "paid_rent"
	BoughtProperty          EventType = "bought_property"
	DeclinedBuy             EventType = "declined_buy"
	BoughtHouse             EventType = "bought_house"
	LandedOnGo              EventType = "landed_on_go"
	PaidTax                 EventType = "paid_tax"
	LandedOnTile            EventType = "landed_on_tile"
	LandedOnUnownedProperty EventType = "landed_on_unowned_property"
	LandedOnOwnProperty     EventType = "landed_on_own_property"
	LandedOnFreeParking     EventType = "landed_on_free_parking"
	LandedOnGoToJail        EventType = "landed_on_go_to_jail"
	LandedOnJail            EventType = "landed_on_jail"
	LandedOnChance          EventType = "landed_on_chance"
	LandedOnCommunityChest  EventType = "landed_on_community_chest"
	LandedOnTax             EventType = "landed_on_tax"
	RolledDice              EventType = "rolled_dice"
	LandedOnStreet          EventType = "landed_on_street"

	InputBuyProperty       EventType = "buy_property"
	InputDeclineBuy        EventType = "decline_buy"
	InputBuyHouse          EventType = "buy_house"
	InputDeclineBuyHouse   EventType = "decline_buy_house"
	InputRollDice          EventType = "roll_dice"
	InputEndTurn           EventType = "end_turn"
	InputPromptNewPlayer   EventType = "prompt_new_player"
	InputPromptGameStart   EventType = "prompt_game_start"
	InputPromptRollDice    EventType = "prompt_roll_dice"
	InputPromptBuyProperty EventType = "prompt_buy_property"
	InputPromptOptions     EventType = "prompt_options"

	CantAfford          EventType = "cant_afford"
	UpdateMoney         EventType = "update_money"
	Jailed              EventType = "jailed"
	ForcedPayToExitJail EventType = "forced_pay_to_exit_jail"
	UpdateProperties    EventType = "update_properties"
	MortgageProperty    EventType = "mortgage_property"
	UnMortgageProperty  EventType = "un_mortgage_property"
)

type GameEvent struct {
	Type    EventType
	Payload any
}

func GetTilePosByName(name string, tiles []Tile) (int, bool) {
	for i, t := range tiles {
		if t.GetName() == name {
			return i, true
		}

	}
	return 0, false
}
