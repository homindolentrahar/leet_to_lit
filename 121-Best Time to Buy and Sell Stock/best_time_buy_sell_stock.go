package main

import (
	"fmt"
	"strconv"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Output: " + strconv.Itoa(maxProfit(prices)))
}

func maxProfit(prices []int) int {
	//
	// First attempt 27 June 2024 14:17
	// Runtime: 88ms, Memory: 8.60MB
	// Time Complexity: O(n), Space Complexity: O(1)
	// Man, I nailed it in the first attempt! (Continue looking for optimal solution)
	//
	//buyIndex := 0
	//profit := 0
	//
	//for i := 1; i < len(prices); i++ {
	//	delta := prices[i] - prices[buyIndex]
	//
	//	if delta > profit {
	//		profit = delta
	//	}
	//	if delta < 0 {
	//		buyIndex = i
	//	}
	//}
	//
	//return profit
	//
	// Second attempt 27 June 2024 15:00
	// Runtime: 123ms, Memory: 12.28MB
	// Time Complexity: O(n), Space Complexity: O(1)
	//
	profit := 0
	buyPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] > buyPrice {
			profit = max(profit, prices[i]-buyPrice)
		} else {
			buyPrice = prices[i]
		}
	}

	return profit
}
