package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println("Output: ", majorityElement(nums))
}

func majorityElement(nums []int) int {
	//
	// First attempt 01 July 2024 15:03
	// Runtime: 25ms, Memory: 6.91MB
	// Time complexity: O(n), Space complexity: O(1)
	// Boyer-Moore Voting Algorithm
	//
	//candidate := -1
	//vote := 0
	//
	//for index := range nums {
	//	if vote == 0 {
	//		candidate = nums[index]
	//		vote++
	//	} else {
	//		if candidate == nums[index] {
	//			vote++
	//		} else {
	//			vote--
	//		}
	//	}
	//}
	//
	//count := 0
	//for index := range nums {
	//	if nums[index] == candidate {
	//		count++
	//	}
	//}
	//
	//if count > len(nums)/2 {
	//	return candidate
	//} else {
	//	return -1
	//}
	//
	// Second attempt 01 July 2024 15:50
	// Runtime: 22ms, Memory: 6.59MB
	// Time complexity: O(n), Space complexity: O(1)
	// Hash-Map approach (NOT RECOMMENDED)
	//
	//hashTable := make(map[int]int)
	//candidate, maxCount := 0, 0
	//
	//for _, element := range nums {
	//	hashTable[element] = hashTable[element] + 1
	//	if hashTable[element] > maxCount {
	//		candidate = element
	//	}
	//	maxCount = max(maxCount, hashTable[element])
	//}
	//
	//return candidate
	//
	// Third attempt 01 July 2024 15:56
	// Runtime: 19ms, Memory: 6.85MB
	// Time complexity: O(n), Space complexity: O(1)
	// Boyer-Moore Voting Algorithm
	//
	candidate, count := 0, 0

	for index := range nums {
		if count == 0 {
			candidate = nums[index]
		}

		if candidate == nums[index] {
			count++
		} else {
			count--
		}
	}

	return candidate
}
