package events

type EventType string

const (
	EventPaidRent                EventType = "paid_rent"
	EventBoughtProperty          EventType = "bought_property"
	EventDeclinedBuy             EventType = "declined_buy"
	EventBoughtHouse             EventType = "bought_house"
	EventLandedOnGo              EventType = "landed_on_go"
	EventPaidTax                 EventType = "paid_tax"
	EventLandedOnUnownedProperty EventType = "landed_on_unowned_property"
	EventLandedOnOwnProperty     EventType = "landed_on_own_property"
	EventLandedOnFreeParking     EventType = "landed_on_free_parking"
	EventLandedOnGoToJail        EventType = "landed_on_go_to_jail"
	EventLandedOnJail            EventType = "landed_on_jail"
	EventLandedOnChance          EventType = "landed_on_chance"
	EventLandedOnCommunityChest  EventType = "landed_on_community_chest"
	EventRolledDice              EventType = "rolled_dice"
)

type GameEvent struct {
	Type    EventType
	Payload any
}

type EventInput string

const (
	InputBuyProperty     EventInput = "buy_property"
	InputDeclineBuy      EventInput = "decline_buy"
	InputBuyHouse        EventInput = "buy_house"
	InputDeclineBuyHouse EventInput = "decline_buy_house"
	InputRollDice        EventInput = "roll_dice"
	InputEndTurn         EventInput = "end_turn"
	InputPromtNewPlayer  EventInput = "prompt_new_player"
	InputPromtGameStart  EventInput = "prompt_game_start"
)

type GameInput struct {
	Type       EventInput
	PlayerName string
	TileName   string
	Amount     int
	Details    string
}

type PaidRentPayload struct {
	PlayerName string
	Owner      string
	TileName   string
	Amount     int
	Details    string
}

type BoughtPropertyPayload struct {
	PlayerName string
	TileName   string
	Amount     int
	Details    string
}

type DeclinedBuyPayload struct {
	PlayerName string
	TileName   string
	Amount     int
}

type LandedOnUnownedPropertyPayload struct {
	PlayerName string
	TileName   string
	Amount     int
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
	Amount     int
}

type LandedOnFreeParkingPayload struct {
	PlayerName string
}

type RolledDicePayload struct {
	PlayerName string
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
