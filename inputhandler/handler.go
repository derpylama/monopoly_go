package inputhandler

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PlayerWantsToBuyProperty(playerName string, propertyName string, price int) bool {
	// Placeholder implementation
	// In a real implementation, this would handle user input to decide
	reader := bufio.NewReader(os.Stdin)
	println(playerName + ", do you want to buy " + propertyName + " for $" + strconv.Itoa(price) + "? (y/n)")

	input, _ := reader.ReadString('\n')
	if input[0] == 'y' || input[0] == 'Y' {
		return true
	} else {
		return false
	}
}

func PlayerEnterNumber(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	println(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		println("Invalid input. Please enter a valid number.")
		return PlayerEnterNumber(prompt)
	}
	return number
}

func PlayerTurnInteraction(playerName string) string {
	reader := bufio.NewReader(os.Stdin)
	println(playerName + ", enter your action (roll, buy, build, mortgage, unmortgage, trade, end):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return strings.ToLower(input)
}

func PlayerWantsToBuildHouse(playerName string, propertyName string, housePrice int) bool {
	// Placeholder implementation
	// In a real implementation, this would handle user input to decide
	return true
}

func PlayerWantsToMortgageProperty(playerName string, propertyName string, mortgageValue int) bool {
	// Placeholder implementation
	// In a real implementation, this would handle user input to decide
	return true
}

func PlayerWantsToUnmortgageProperty(playerName string, propertyName string, unmortgageCost int) bool {
	// Placeholder implementation
	// In a real implementation, this would handle user input to decide
	return true
}

func PlayerWantsToTradeProperty(playerName string, propertyName string, tradeWith string) bool {
	// Placeholder implementation
	// In a real implementation, this would handle user input to decide
	return true
}

func PlayerWantsToEndTurn(playerName string) bool {
	// Placeholder implementation
	// In a real implementation, this would handle user input to decide
	return true
}
