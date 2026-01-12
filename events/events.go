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
)

type GameEvent struct {
	Type       EventType
	PlayerName string
	Title      string
	Owner      string
	TileName   string
	Amount     int
	Details    string
}

type EventInput string

const (
	InputBuyProperty     EventInput = "buy_property"
	InputDeclineBuy      EventInput = "decline_buy"
	InputBuyHouse        EventInput = "buy_house"
	InputDeclineBuyHouse EventInput = "decline_buy_house"
	InputRollDice        EventInput = "roll_dice"
)

type GameInput struct {
	Type       EventInput
	PlayerName string
	TileName   string
	Amount     int
	Details    string
}
