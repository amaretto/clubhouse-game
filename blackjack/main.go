package main

import (
	"fmt"
	"strconv"
	"time"
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

type Player struct {
	name   string
	hand   []int
	result int
	chip   int
	com    bool
}

type Dealer struct {
	Player
	game *Game
}

func main() {
	g := Game{}
	g.dealer = Dealer{Player{name: "Dealer"}, &g}
	g.players = []Player{Player{name: "A"}}
	g.roundNum = roundNum
	g.deck = createDeck(deckNum, shuffleTime)

	g.playGame()
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

func (d *Dealer) dealCards() {
	d.hand = []int{}
	for i := 0; i < len(d.game.players); i++ {
		d.game.players[i].hand = []int{}
	}
	for i := 0; i < 2; i++ {
		d.hand = append(d.hand, d.game.deck.draw())
		for j := 0; j < len(d.game.players); j++ {
			d.game.players[j].hand = append(d.game.players[j].hand, d.game.deck.draw())
		}
	}
}

func (d *Dealer) showDealerHands() {
	fmt.Printf("\x1b[31m")
	fmt.Println("\n///// Dealer Hands /////")
	fmt.Printf("Cards: %d, *\n", d.hand[0])
	fmt.Printf("\x1b[0m")
}

func (d *Dealer) processPlayers() {
	var input string
	var newCard int
	var p *Player
	for i := 0; i < len(d.game.players); i++ {
		p = &d.game.players[i]
		fmt.Printf("\x1b[34m")
		fmt.Printf("\n\n////// Player %s Turn!! //////\n", p.name)
		for {
			fmt.Printf("Player %s Hand: %s\n", p.name, joinHands(p.hand))

			pj := judgeHand(countTotal(p.hand))
			if pj == 0 {
				fmt.Println("Player Total: Blackjack!!!")
				p.result = 21
				fmt.Printf("\x1b[0m")
				break
			} else if pj == 1 {
				fmt.Println("Player Total:", countTotal(p.hand))
			} else {
				fmt.Println("Player Total: Bursted...")
				p.result = minCount(countTotal(p.hand))
				fmt.Printf("\x1b[0m")
				break
			}

			fmt.Println("Hit(h)/Stand(s)?")
			fmt.Scan(&input)

			if input == "h" {
				newCard = d.game.deck.draw()
				fmt.Println("New Card: ", newCard)
				p.hand = append(p.hand, newCard)
			} else if input == "s" {
				p.result = maxAvailableCount(countTotal(p.hand))
				fmt.Println("\nPlayer Total:", p.result)
				fmt.Printf("\x1b[0m")
				break
			} else {
				fmt.Println("invalid input")
				continue
			}
		}
	}
}

func (d *Dealer) processDealer() {
	var newCard int
	fmt.Printf("\x1b[31m")
	fmt.Println("\n\n////// Dealer Turn! //////")
	for {
		fmt.Println("Delaer Hand:", joinHands(d.hand))
		delay()

		dj := judgeHand(countTotal(d.hand))
		if dj == 0 {
			fmt.Println("Dealer Black Jack!!")
			fmt.Printf("\x1b[0m")
			d.result = 21
			return
		} else if dj == 1 {
			currentMax := maxAvailableCount(countTotal(d.hand))
			if currentMax <= 21 && currentMax >= 17 {
				d.result = currentMax
				fmt.Println("\nDelaer Total:", d.result)
				fmt.Printf("\x1b[0m")
				return
			}
			fmt.Println("\nDelaer Total:", countTotal(d.hand))
		} else {
			fmt.Println("Dealer Bursted..")
			fmt.Printf("\x1b[0m")
			d.result = minCount(countTotal(d.hand))
			return
		}
		newCard = d.game.deck.draw()
		fmt.Println("New Card: ", newCard)
		d.hand = append(d.hand, newCard)
	}
}

func (d *Dealer) judgeResults() {
	fmt.Printf("\n\n//////   Judge!   //////\n")
	delay()

	fmt.Printf("\nDealer Total:%d", d.result)

	var p Player
	for i := 0; i < len(d.game.players); i++ {
		delay()
		p = d.game.players[i]
		fmt.Printf("\n\nPlayer %s Total: %d\n", p.name, p.result)
		delay()
		if p.result > 21 && d.result > 21 || p.result == d.result {
			fmt.Printf("\nPlayer %s Draw\n", p.name)
		} else if p.result <= 21 && (p.result > d.result || d.result > 21) {
			fmt.Printf("\nPlayer %s Win!\n", p.name)
			// ToDo: process chip
		} else {
			fmt.Printf("\nPlayer %s Lose...\n", p.name)
			// ToDo: process chip
		}
	}
	delay()
}

func maxAvailableCount(counts []int) int {
	max := 0
	for _, c := range counts {
		if c > max && 21 >= c {
			max = c
		}
	}
	return max
}

func minCount(counts []int) int {
	min := 100
	for _, c := range counts {
		if c < min {
			min = c
		}
	}
	return min
}

func judgeResult(p Player, d Dealer) {
	time.Sleep(1 * time.Second)
	fmt.Println("\n\n//////   Judge!   //////")
	time.Sleep(1 * time.Second)
	fmt.Printf("Player:%d, Dealer:%d\n\n", p.result, d.result)
	time.Sleep(1 * time.Second)
	if p.result > 21 && d.result > 21 || p.result == d.result {
		fmt.Println("Draw")
	} else if p.result <= 21 && (p.result > d.result || d.result > 21) {
		fmt.Println("Player Win!")
	} else {
		fmt.Println("Dealer Win!")
	}
}

// 0:blackjuck, 1:continue, 2:burst
func judgeHand(counts []int) int {
	burstCount := 0
	for _, c := range counts {
		if c > 21 {
			burstCount++
		} else if c == 21 {
			return 0
		}
	}
	if burstCount == len(counts) {
		return 2
	}
	return 1
}

func countTotal(hands []int) []int {
	var result []int
	if hands[0] == 1 {
		result = []int{1, 11}
	} else if hands[0] < 10 {
		result = []int{hands[0]}
	} else {
		result = append(result, 10)
	}

	for i := 1; i < len(hands); i++ {
		if hands[i] == 1 {
			tmp := []int{}
			for _, r := range result {
				tmp = append(tmp, r+1)
				tmp = append(tmp, r+11)
			}
			result = tmp
		} else {
			var delta int
			if hands[i] < 10 {
				delta = hands[i]
			} else {
				delta = 10
			}
			for j := 0; j < len(result); j++ {
				result[j] += delta
			}
		}
	}
	return result
}

func delay() {
	time.Sleep(1 * time.Second)
}

func joinHands(cards []int) string {
	var result string
	for i := 0; i < len(cards); i++ {
		var tmp string
		if cards[i] < 11 {
			tmp = strconv.Itoa(cards[i])
		} else {
			if cards[i] == 11 {
				tmp = "J"
			} else if cards[i] == 12 {
				tmp = "Q"
			} else {
				tmp = "K"
			}
		}
		if result == "" {
			result = tmp
		} else {
			result += ", " + tmp
		}
	}
	return result
}
