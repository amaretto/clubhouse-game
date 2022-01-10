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
	name   string
	hand   []int
	result int
	chip   int
}

type Dealer struct {
	Player
}

type Deck struct {
	n     int
	cards []int
}

type Game struct {
	round int
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
	fmt.Printf("\x1b[32m")
	fmt.Printf("///////////////////////////////\n")
	fmt.Printf("//////////start game!//////////\n")
	fmt.Printf("///////////////////////////////\n")
	fmt.Printf("\x1b[0m")

	fmt.Printf("\x1b[31m")
	fmt.Println("\n/////Dealer Hand/////")
	fmt.Printf("Cards: %d,*\n", d.hand[0])
	fmt.Printf("\x1b[0m")

	fmt.Printf("\x1b[34m")
	fmt.Println("\n//////Your Hand//////")
	fmt.Println("Cards: " + joinHands(p.hand))
	fmt.Printf("Total: %d\n\n", countTotal(p.hand))

	// player turn
	var input string
	fmt.Println("\n//////Your Turn!!//////\n")
	for {
		fmt.Println("Hit(h)/Stand(s)?")
		fmt.Scan(&input)
		if input == "h" {
			fmt.Println("New Card: ", cards[n])
			p.hand = append(p.hand, cards[n])
			n++
			fmt.Println("Your Hands: ", joinHands(p.hand))
			// do judge
			pj := judgeHand(countTotal(p.hand))
			if pj == 0 {
				fmt.Println("Player Total: Blackjack!!!")
				p.result = 21
				break
			} else if pj == 1 {
				fmt.Println("Player Total:", countTotal(p.hand))
				continue
			} else {
				fmt.Println("Player Total: Bursted...")
				p.result = minCount(countTotal(p.hand))
				break
			}
		} else if input == "s" {
			p.result = maxCount(countTotal(p.hand))
			fmt.Println("\nPlayer Total:", p.result)
			break
		} else {
			fmt.Println("invalid input")
			continue
		}
	}
	fmt.Printf("\x1b[0m")

	// dealer turn
	dealerProcess(cards, &d)

	judgeResult(p, d)
}

func playerProcess(cards []int, p Player) {
	fmt.Println("\n\n------Player Turn------")

}

func dealerProcess(cards []int, d *Player) {
	fmt.Printf("\x1b[31m")
	fmt.Println("\n\n//////Dealer Turn!//////")

	for {
		fmt.Println("Delaer Hand:", joinHands(d.hand))
		time.Sleep(1 * time.Second)
		currentMax := maxCount(countTotal(d.hand))
		if currentMax >= 17 {
			d.result = currentMax
			fmt.Println("Dealer Total:", d.result)
			fmt.Printf("\x1b[0m")
			return
		}
		fmt.Println("New Card: ", cards[n])
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
	time.Sleep(1 * time.Second)
	fmt.Println("\n\n//////   Judge!   //////\n")
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
