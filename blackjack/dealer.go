package main

import "fmt"

type Dealer struct {
	Player
	game *Game
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