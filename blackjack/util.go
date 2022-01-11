package main

import (
	"strconv"
	"time"
)

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
