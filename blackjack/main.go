package main

import (
	"fmt"
)

const (
	// game
	roundNum  = 3
	playerNum = 1
	// deck
	deckNum     = 4
	shuffleTime = 100
)

type Game struct {
	dealer   Dealer
	players  []Player
	roundNum int
	deck     Deck
}

func (g *Game) playGame() {
	fmt.Printf("\x1b[32m")
	fmt.Printf("/////////////////////////////////\n")
	fmt.Printf("////////// game start! //////////\n")
	fmt.Printf("/////////////////////////////////\n")
	fmt.Printf("\x1b[0m")

	delay()

	for i := 0; i < g.roundNum; i++ {
		fmt.Printf("\x1b[32m")
		fmt.Printf("\n\n----- Round %d -----\n", i+1)
		fmt.Printf("\x1b[0m")
		g.deck.shuffle()
		g.round()
	}
	// ToDo: judge total results
}

func (g *Game) round() {
	delay()
	g.dealer.dealCards()
	g.dealer.showDealerHands()
	delay()
	g.dealer.processPlayers()
	delay()
	g.dealer.processDealer()
	delay()
	g.dealer.judgeResults()
}

func main() {
	g := Game{}
	g.dealer = Dealer{Player{name: "Dealer"}, &g}
	g.players = []Player{Player{name: "A"}}

	g.roundNum = roundNum
	g.deck = createDeck(deckNum, shuffleTime)

	g.playGame()
}
