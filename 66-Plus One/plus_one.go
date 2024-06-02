package main

import "fmt"

func main() {
	nums := []int{9}
	fmt.Println("Result: ", plusOne(nums))
}

func plusOne(digits []int) []int {
	//
	// First attempt Jun 02 2024 12:54
	// Runtime: 1ms, Memory: 2.28MB
	//
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
