package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("Output: ", containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	//
	// First attempt 02 July 2024 10:04
	// Runtime: 85ms, Memory: 13.34MB
	// Time complexity: O(n), Space complexity: O(n)
	// HashMap
	//
	occurrenceMap := make(map[int]int)
	for index := range nums {
		if _, ok := occurrenceMap[nums[index]]; ok {
			return true
		}
		occurrenceMap[nums[index]] += 1
	}

	return false
	//
	// Second attempt 02 July 2024 11:04
	// Runtime: 134ms, Memory: 8.76MB
	// Time complexity: O(n), Space complexity: O(n)
	// HashMap
	//
	//occurrenceMap := make(map[int]bool)
	//for index := range nums {
	//	if _, ok := occurrenceMap[nums[index]]; ok {
	//		return true
	//	}
	//
	//	occurrenceMap[nums[index]] = true
	//}
	//
	//return false
}
