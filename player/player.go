package player

import (
	cards "monopoly/card"
)

type Player struct {
	Position    int
	Money       int
	Name        string
	cardsInHand []cards.Card
	isJailed    bool
}
