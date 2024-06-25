package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}

	fmt.Println("Output: ", singleNumber(nums))
}

func singleNumber(nums []int) int {
	//
	// First attempt 25 June 2024 14:53
	// Runtime: 13ms, Memory: 6.18
	//
	//if len(nums) < 2 {
	//	return nums[0]
	//}
	//
	//val := 0
	//
	//for i := 0; i < len(nums); i++ {
	//	val = val ^ nums[i]
	//}
	//
	//return val
	//
	// Second attempt 25 June 2024 15:57
	// Runtime: 14ms, Memory: 6.13MB
	//
	val := 0
	for _, num := range nums {
		val = val ^ num
	}

	return val
}
