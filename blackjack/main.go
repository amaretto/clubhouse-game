package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	hand []int
}

const (
	deck        = 2
	shuffleTime = 100
)

func main() {
	cnum := deck * 13

	cards := make([]int, cnum)

	for i := 0; i < cnum; i++ {
		cards[i] = i%13 + 1
	}
	fmt.Println(cards)
	shuffle(cards, cnum)
}

func shuffle(cards []int, cnum int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < shuffleTime; i++ {
		a := rand.Intn(cnum)
		b := rand.Intn(cnum)
		cards[a], cards[b] = cards[b], cards[a]
	}
}
