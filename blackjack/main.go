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
	fmt.Printf("////////// GAME START! //////////\n")
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

	var p *Player
	for i := 0; i < len(g.players); i++ {
		p = &g.players[i]
		if p.chip > 1000 {
			fmt.Printf("Player %s won $%d\n", p.name, p.chip-1000)
		} else if p.chip == 1000 {
			fmt.Printf("Player %s chip don't change\n", p.name)
		} else {
			fmt.Printf("Player %s lose $%d\n", p.name, 1000-p.chip)
		}
	}
}

func (g *Game) round() {
	delay()
	g.dealer.confirmBet()
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
	g.dealer = Dealer{Player{name: "Dealer", chip: 1000000}, &g}
	g.players = []Player{Player{name: "A", chip: 1000}}

	g.roundNum = roundNum
	g.deck = createDeck(deckNum, shuffleTime)

	g.playGame()
}
