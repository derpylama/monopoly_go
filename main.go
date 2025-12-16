package main

import (
	"fmt"
	board "monopoly/board"
)

func main() {
	fmt.Println("t")
	//fmt.Scan(&test)

	//fmt.Println(test)
	board := board.NewBoard()

	for _, element := range board.Tiles() {
		fmt.Print(element.GetName() + " \n")
	}
}
