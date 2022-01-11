package main

import (
	"strconv"
	"time"
)

func maxAvailableSum(sums []int) int {
	max := 0
	for _, sum := range sums {
		if sum > max && 21 >= sum {
			max = sum
		}
	}
	return max
}

func minOverSum(counts []int) int {
	min := 100
	for _, c := range counts {
		if c < min {
			min = c
		}
	}
	return min
}

func calcSums(hands []int) []int {
	sums := []int{0}
	var d int

	for i := 0; i < len(hands); i++ {
		if hands[i] == 1 {
			tmp := []int{}
			for _, r := range sums {
				tmp = append(tmp, r+1)
				tmp = append(tmp, r+11)
			}
			sums = tmp
		} else {
			if hands[i] < 10 {
				d = hands[i]
			} else {
				d = 10
			}
			for j := 0; j < len(sums); j++ {
				sums[j] += d
			}
		}
	}
	return sums
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
func checkSums(sums []int) int {
	burstCount := 0
	for _, c := range sums {
		if c > 21 {
			burstCount++
		} else if c == 21 {
			return 0
		}
	}
	if burstCount == len(sums) {
		return 2
	}
	return 1
}
