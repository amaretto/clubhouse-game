package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	deckNum     = 4
	shuffleTime = 100
)

var (
	n int
)

type Player struct {
	hand   []int
	result int
	chip   int
}

type Deck struct {
	n     int
	cards []int
}

func main() {
	cnum := deckNum * 13
	cards := make([]int, cnum)

	for i := 0; i < cnum; i++ {
		cards[i] = i%13 + 1
	}
	shuffle(cards, cnum)

	playGame(cards)
}

func shuffle(cards []int, cnum int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < shuffleTime; i++ {
		a := rand.Intn(cnum)
		b := rand.Intn(cnum)
		cards[a], cards[b] = cards[b], cards[a]
	}
}

func playGame(cards []int) {
	var d, p Player

	// ToDo: adopt multiple player
	p.hand = append(p.hand, cards[0])
	d.hand = append(d.hand, cards[1])
	p.hand = append(p.hand, cards[2])
	d.hand = append(d.hand, cards[3])
	n = 4

	// notify info
	fmt.Println("\n\n/////Dealer/////")
	fmt.Printf("%d,*\n", d.hand[0])

	fmt.Println("/////You/////")
	fmt.Println(join(p.hand))
	fmt.Printf("Total : %d\n\n", countTotal(p.hand))

	// player turn
	for {
		fmt.Println("Hit(h)/Stand(s)?")
		var input string
		fmt.Scan(&input)
		if input == "h" {
			fmt.Println("card: ", cards[n])
			p.hand = append(p.hand, cards[n])
			n++
			// do judge
			pj := judgeHand(countTotal(p.hand))
			if pj == 0 {
				fmt.Println("Blackjack!")
				p.result = 21
				break
			} else if pj == 1 {
				continue
			} else {
				fmt.Println("Burst!!")
				p.result = minCount(countTotal(p.hand))
				break
			}
		} else if input == "s" {
			p.result = maxCount(countTotal(p.hand))
			break
		} else {
			fmt.Println("invalid input")
			continue
		}
	}

	// dealer turn
	dealerProcess(cards, &d)

	judgeResult(p, d)
}

func playerProcess(cards []int, p Player) {
	fmt.Println("\n\n------Player Turn------")

}

func dealerProcess(cards []int, d *Player) {
	fmt.Println("\n\n------Delaler Turn------")

	for {
		currentMax := maxCount(countTotal(d.hand))
		if currentMax >= 17 {
			d.result = currentMax
			return
		}
		d.hand = append(d.hand, cards[n])
		n++
	}
}

func printPlayerInfo(p Player) {

}

func maxCount(counts []int) int {
	max := 0
	for _, c := range counts {
		if c > max {
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

func judgeResult(p, d Player) {
	fmt.Printf("Player:%d, Delaer:%d\n", p.result, d.result)
	if p.result > 21 && d.result > 21 || p.result == d.result {
		fmt.Println("Draw")
	} else if p.result <= 21 && (p.result > d.result || d.result > 21) {
		fmt.Println("Player Win!")
	} else {
		fmt.Println("Delaer Win!")
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
	fmt.Println("Total:", join(result))
	return result
}

func join(cards []int) string {
	result := strconv.Itoa(cards[0])
	for i := 1; i < len(cards); i++ {
		result += ", " + strconv.Itoa(cards[i])
	}
	return result
}
