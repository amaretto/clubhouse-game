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

func calcSums(hands []Card) []int {
	sums := []int{0}
	var d int

	for i := 0; i < len(hands); i++ {
		if hands[i].num == 1 {
			tmp := []int{}
			for _, r := range sums {
				tmp = append(tmp, r+1)
				tmp = append(tmp, r+11)
			}
			sums = tmp
		} else {
			if hands[i].num < 10 {
				d = hands[i].num
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

func convCardNum(cardNum int) string {
	var result string
	if cardNum < 11 {
		result = strconv.Itoa(cardNum)
	} else if cardNum == 11 {
		result = "J"
	} else if cardNum == 12 {
		result = "Q"
	} else {
		result = "K"
	}
	return result
}

func joinHands(cards []Card) string {
	var result string
	for i := 0; i < len(cards); i++ {
		if result == "" {
			result = convCardNum(cards[i].num)
		} else {
			result += ", " + convCardNum(cards[i].num)
		}
	}
	return result
}

func joinSums(sums []int) string {
	var result string
	for i := 0; i < len(sums); i++ {
		if result == "" {
			result = strconv.Itoa(sums[i])
		} else {
			result += "," + strconv.Itoa(sums[i])
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
