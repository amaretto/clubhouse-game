package main

import "fmt"

type Player struct {
	hand []int
}

var (
	deck = 2
)

func main() {
	cards := make([]int, deck*13)
	for i := 0; i < 13; i++ {
		cards[i] = i + 1
	}
	fmt.Println(cards)
}

func shuffleCards(cards *[]int) {
}
