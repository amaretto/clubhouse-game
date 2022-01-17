package main

import (
	"math/rand"
	"time"
)

var (
	suits = map[int]string{0: "spade", 1: "club", 2: "diamond", 3: "heart"}
)

type Deck struct {
	pos         int
	shuffleTime int
	cards       []Card
}

type Card struct {
	num  int
	suit int //0:spade, 1:club, 2:diamond, 3:heart
}

func createDeck(deckNum, shuffleTime int) Deck {
	deck := Deck{}
	cnum := deckNum * 13 * 4
	var num, suit int

	deck.cards = make([]Card, cnum)
	deck.shuffleTime = shuffleTime
	for i := 0; i < cnum; i++ {
		num = i%13 + 1
		suit = i / 13
		deck.cards[i] = Card{num, suit}
	}
	return deck
}

func (d *Deck) draw() Card {
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
