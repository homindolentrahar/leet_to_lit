package main

import "fmt"

func main() {
	rowIndex := 4
	fmt.Println(getRow(rowIndex))
}

func getRow(rowIndex int) []int {
	//
	// First attempt 26 June 2024 11:14
	// Runtime: 1ms, Memory: 2.40MB
	// Recursive function
	// Time complexity: O(n^2), Space complexity: O(n)
	//
	//if rowIndex == 0 {
	//	return []int{1}
	//}
	//
	//previousRow := getRow(rowIndex - 1)
	//var row []int
	//
	//row = append(row, 1)
	//for i := 1; i < rowIndex; i++ {
	//	row = append(row, previousRow[i-1]+previousRow[i])
	//}
	//row = append(row, 1)
	//
	//return row
	//
	// Second attempt 26 June 2024
	// Runtime: 1ms, Memory: 2.32MB
	// Nested loop
	// Time complexity: O(n^2), Space complexity: O(n)
	//
	//res := []int{1}
	//
	//for range rowIndex {
	//	row := make([]int, len(res)+1)
	//
	//	for j := range len(res) {
	//		row[j] += res[j]
	//		row[j+1] += res[j]
	//	}
	//
	//	res = row
	//}
	//
	//return res
	//
	// Third attempt 26 June 2024 11:47
	// Runtime: 1ms, Memory: 2.19MB
	// For loop
	// Time complexity: O(n), Space complexity: O(n)
	//
	res := []int{1}
	prev := 1

	for i := 1; i <= rowIndex; i++ {
		val := prev * (rowIndex - i + 1) / i
		res = append(res, val)
		prev = val
	}

	return res
}
