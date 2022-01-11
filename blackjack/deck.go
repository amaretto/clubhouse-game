package main

import (
	"math/rand"
	"time"
)

type Deck struct {
	pos         int
	shuffleTime int
	cards       []int
}

func createDeck(deckNum, shuffleTime int) Deck {
	deck := Deck{}
	cnum := deckNum * 13 * 4

	deck.cards = make([]int, cnum)
	deck.shuffleTime = shuffleTime

	for i := 0; i < cnum; i++ {
		deck.cards[i] = i%13 + 1
	}
	return deck
}

func (d *Deck) draw() int {
	newCard := d.cards[d.pos]
	d.pos++
	return newCard
}

func (d *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < d.shuffleTime; i++ {
		a := rand.Intn(len(d.cards))
		b := rand.Intn(len(d.cards))
		d.cards[a], d.cards[b] = d.cards[b], d.cards[a]
	}
	d.pos = 0
}
