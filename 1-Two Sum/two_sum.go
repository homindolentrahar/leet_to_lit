package main

import "fmt"

func main() {
	target := 6
	nums := []int{3, 2, 4}

	fmt.Println("Result: ", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// First attempt Apr 16, 2024 20:26
	// Runtime 6ms, Memory 4.17MB
	//
	mapResult := make(map[int]int)

	for index, num := range nums {
		if _, ok := mapResult[num]; ok {
			return []int{mapResult[num], index}
		}

		diff := target - num
		mapResult[diff] = index
	}

	return []int{}
	//
	// Second attempt Apr 16, 2024 20:46
	// Runtime 8ms, Memory 5.49MB
	//
	//	mapDiffs := make(map[int]int)
	//	var indexes []int
	//
	//	for index, num := range nums {
	//		diff := target - num
	//
	//		if _, ok := mapDiffs[diff]; ok {
	//			indexes = []int{mapDiffs[diff], index}
	//		}
	//
	//			mapDiffs[num] = index
	//		}
	//
	//		return indexes
	//
	// Third attempt Apr 16, 2024 21:00
	// Runtime 7ms, Memory 5.52MB
	//	mapDiffs := make(map[int]int)
	//	var indexes []int
	//
	//	for index, num := range nums {
	//		diff := target - num
	//
	//		if n, ok := mapDiffs[diff]; ok {
	//			indexes = []int{n, index}
	//		}
	//
	//		mapDiffs[num] = index
	//	}
	//
	//	return indexes
}
