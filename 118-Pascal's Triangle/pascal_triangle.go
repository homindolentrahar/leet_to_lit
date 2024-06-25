package main

import "fmt"

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	//
	// First attempt 24 June 2024 15:28
	// Runtime: 1ms, Memory: 2.72MB
	//
	//res := [][]int{{1}}
	//
	//for range numRows - 1 {
	//	lastRest := res[len(res)-1]
	//	var temp = []int{0}
	//	var row []int
	//
	//	temp = append(temp, lastRest...)
	//	temp = append(temp, 0)
	//
	//	for j := range len(lastRest) + 1 {
	//		row = append(row, temp[j]+temp[j+1])
	//	}
	//
	//	res = append(res, row)
	//}
	//
	//return res
	//
	// Second attempt 24 June 2024 15:56
	// Runtime: 1ms, Memory: 2.63MB
	//
	//res := make([][]int, numRows)
	//
	//for i := 1; i <= numRows; i++ {
	//	solve := 1
	//	row := make([]int, i)
	//	row[0] = solve
	//
	//	for j := 1; j < i; j++ {
	//		solve = (solve * (i - j)) / j
	//		row[j] = solve
	//	}
	//
	//	res[i-1] = row
	//}
	//
	//return res
	//
	// Third attempt 24 June 2024 16:20
	// Runtime: 1ms, Memory: 2.63MB
	// Combinatorial Formula
	//
	//res := [][]int{{1}}
	//
	//for i := 1; i < numRows; i++ {
	//	prevRow := res[i-1]
	//	var row []int
	//	row = append(row, 1)
	//
	//	for j := 1; j < i; j++ {
	//		row = append(row, prevRow[j-1]+prevRow[j])
	//	}
	//
	//	row = append(row, 1)
	//	res = append(res, row)
	//}
	//
	//return res
	//
	// Fourth Attempt 24 June 2024 16:42
	// Runtime: 1ms, Memory: 2.58MB
	// Dynamic Programming with 1D Array
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}

	res := generate(numRows - 1)
	prevRow := res[numRows-2]
	var row []int

	row = append(row, 1)
	for i := 1; i < numRows-1; i++ {
		row = append(row, prevRow[i-1]+prevRow[i])
	}
	row = append(row, 1)
	res = append(res, row)

	return res
}
