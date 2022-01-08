package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Player struct {
	hand []int
}

const (
	deck        = 4
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
	fmt.Println(cards)
	game(cards)
}

func shuffle(cards []int, cnum int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < shuffleTime; i++ {
		a := rand.Intn(cnum)
		b := rand.Intn(cnum)
		cards[a], cards[b] = cards[b], cards[a]
	}
}

func game(cards []int) {
	var d, p Player
	p.hand = append(p.hand, cards[0])
	d.hand = append(d.hand, cards[1])
	p.hand = append(p.hand, cards[2])
	d.hand = append(d.hand, cards[3])
	n := 4

	// notify info
	fmt.Println("\n\n/////Dealer/////")
	fmt.Printf("%d,*\n", d.hand[0])

	fmt.Println("/////You/////")
	fmt.Println(join(p.hand))
	fmt.Printf("Total : %d\n\n", count(p.hand))

	// player turn
	pj := 0
	for {
		fmt.Println("Hit(h)/Stand(s)?")
		var input string
		fmt.Scan(&input)
		if input == "h" {
			fmt.Println("card: ", cards[n])
			p.hand = append(p.hand, cards[n])
			n++
			// do judge
			pj = judge(count(p.hand))
			if pj == 0 || pj == 2 {
				break
			} else if pj == 1 {
				continue
			}
		} else if input == "s" {
			break
		} else {
			fmt.Println("invalid input")
			continue
		}
	}
	fmt.Println("pj:", pj)

	// dealer turn
}

// 0:blackjuck, 1:continue, 2:burst
func judge(counts []int) int {
	burstCount := 0
	for c := range counts {
		if c > 21 {
			burstCount++
		} else if c == 21 {
			return 0
		}
	}
	if burstCount == len(counts) {
		return 2
	}
	fmt.Println("continue")
	return 1
}

func count(hands []int) []int {
	var result []int
	if hands[0] < 10 {
		result = []int{hands[0]}
	} else {
		result = append(result, 10)
	}
	for i := 1; i < len(hands); i++ {
		if hands[i] == 1 {
			tmp := []int{}
			for r := range result {
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
