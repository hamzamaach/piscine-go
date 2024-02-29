package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	players := []string{"Player 1", "Player 2", "Player 3", "Player 4"}
	cardsPerPlayer := 3
	for i, player := range players {
		startIndex := i * cardsPerPlayer
		endIndex := (i + 1) * cardsPerPlayer
		cardForPlayer := deck[startIndex:endIndex]
		fmt.Printf("%s: %d, %d, %d\n", player, cardForPlayer[0], cardForPlayer[1], cardForPlayer[2])
	}
}
