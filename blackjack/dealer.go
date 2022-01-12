package main

import "fmt"

type Dealer struct {
	Player
	game *Game
}

func (d *Dealer) confirmBet() {
	var input string
	for i := 0; i < len(d.game.players); i++ {
		p := d.game.players[i]
		fmt.Printf("Player %s JOIN GAME?(y/n)\n", p.name)
		fmt.Scan(&input)
		if input == "y" {
			p.status = 1
		} else {
			p.status = 2
		}
	}
}

func (d *Dealer) dealCards() {
	d.hand = []int{}
	for i := 0; i < len(d.game.players); i++ {
		d.game.players[i].status = 0
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
	fmt.Printf("DealerHands: %d, *\n", d.hand[0])
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
			if p.status == 0 {
				fmt.Println("Continue the game?: y/n(surrender)")
				fmt.Scan(&input)

				if input == "y" {
					p.status = 1
				} else {
					p.status = 3
					return
				}
			}

			pj := checkSums(calcSums(p.hand))
			if pj == 0 {
				fmt.Println("Player Total: Blackjack!!!")
				p.result = 21
				fmt.Printf("\x1b[0m")
				break
			} else if pj == 1 {
				fmt.Println("Player Total:", calcSums(p.hand))
			} else {
				fmt.Println("Player Total: Bursted...")
				p.result = minOverSum(calcSums(p.hand))
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
				p.result = maxAvailableSum(calcSums(p.hand))
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

		dj := checkSums(calcSums(d.hand))
		if dj == 0 {
			fmt.Println("Dealer Black Jack!!")
			fmt.Printf("\x1b[0m")
			d.result = 21
			return
		} else if dj == 1 {
			currentMax := maxAvailableSum(calcSums(d.hand))
			if currentMax <= 21 && currentMax >= 17 {
				d.result = currentMax
				fmt.Println("\nDelaer Total:", d.result)
				fmt.Printf("\x1b[0m")
				return
			}
			fmt.Println("\nDelaer Total:", calcSums(d.hand))
		} else {
			fmt.Println("Dealer Bursted..")
			fmt.Printf("\x1b[0m")
			d.result = minOverSum(calcSums(d.hand))
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
