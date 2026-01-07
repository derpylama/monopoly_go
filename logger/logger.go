package logger

import (
	"fmt"
	"monopoly/tile"
	"strconv"
)

func LogOnLand(ownerName string, tileName string, isOwned bool, amountToPay int, currentPlayerName string, landedOnTile tile.Tile) {
	ClearScreen()
	switch landedOnTile.(type) {
	case *tile.Street:
		if isOwned {
			println(currentPlayerName + " landed on " + tileName + " owned by " + ownerName + " and pays " + strconv.Itoa(amountToPay))
		} else {
			println(currentPlayerName + " landed on " + tileName + " which is unowned.")
		}
	case *tile.TrainStation:
		if isOwned {
			println(currentPlayerName + " landed on " + tileName + " owned by " + ownerName + " and pays " + strconv.Itoa(amountToPay))
		} else {
			println(currentPlayerName + " landed on " + tileName + " which is unowned.")

		}
	case *tile.TaxTile:
		println(currentPlayerName + " landed on " + tileName + " and pays tax of " + strconv.Itoa(amountToPay))

	case *tile.ChanceTile:
		println(currentPlayerName + " landed on " + tileName)

	case *tile.CommunityChest:
		println(currentPlayerName + " landed on " + tileName)

	case *tile.GoToJail:
		println(currentPlayerName + " landed on Go To Jail Tile and is sent to Jail")

	case *tile.JailTile:
		println(currentPlayerName + " is just visiting Jail")

	default:
		println(currentPlayerName + " landed on " + tileName)
	}

}

func LogRollDice(playerName string, roll []int, currentMoney int) {
	ClearScreen()

	println(playerName + " rolled a " + strconv.Itoa(roll[0]) + " and a " + strconv.Itoa(roll[1]) + " for a total of " + strconv.Itoa(roll[0]+roll[1]))
	println("Current money: " + strconv.Itoa(currentMoney))
}

func LogBuyProperty(playerName string, propertyName string, propertyPrice int, remainingMoney int) {
	println(playerName + " bought " + propertyName + " for " + strconv.Itoa(propertyPrice) + ". Remaining money: " + strconv.Itoa(remainingMoney))
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func LogPlayersProperties(playerName string, ownedProperties []tile.Property) {
	ClearScreen()
	println(playerName + "'s Properties:")

	count := 0

	for _, t := range ownedProperties {

		fmt.Println("\n" + strconv.Itoa(count) + " - " + t.GetName() + " | Price: " + strconv.Itoa(t.GetPrice()) + " | Rent: " + strconv.Itoa(t.GetRent(nil, nil)) + " | Houses: " + strconv.Itoa(t.(*tile.Street).GetHouseAmount()))
		count++
	}
}
