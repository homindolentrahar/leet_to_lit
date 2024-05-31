package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 2}
	fmt.Println("Output: ", removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	//
	// Failed attempt (Cannot return the right value, but the position of the elements are correct)
	//
	//if len(nums) == 0 {
	//	return 0
	//}
	//
	//current := 0
	//lastIndex := len(nums) - 1
	//
	//for current < lastIndex {
	//	if nums[current] == val {
	//		if nums[lastIndex] == val {
	//			lastIndex--
	//		}
	//
	//		temp := nums[current]
	//		nums[current] = nums[lastIndex]
	//		nums[lastIndex] = temp
	//
	//		lastIndex--
	//	}
	//
	//	current++
	//}
	//
	//fmt.Println(nums)
	//
	//if current == 0 {
	//	return 0
	//} else {
	//	return current + 1
	//}
	//
	// Second attempt May 31, 2024 20:58
	// Runtime 1ms, Memory 2.35MB
	//
	pointer := 0
	for index := range nums {
		if nums[index] != val {
			nums[pointer] = nums[index]
			pointer++
		}
	}

	fmt.Println("Nums: ", nums)

	return pointer
}
