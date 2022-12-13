package main

import (
	"fmt"
	"math"
)

var mW int
var mH int
var mp map[uint64]int

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func goUp(m [][]int, x, y int) int {
	if y == 0 && x < mW && x >= 0 {
		return m[y][x]
	}
	if x >= mW || x < 0 {
		return math.MaxInt32
	}
	key := uint64(x)<<32 | uint64(y)
	if v, ok := mp[key]; ok {
		return v
	}

	mp[key] = m[y][x] +
		min(
			min(
				goUp(m, x+1, y-1),
				goUp(m, x, y-1)),
			goUp(m, x-1, y-1))
	return mp[key]
}

func minFallingPathSum(matrix [][]int) int {
	mH, mW = len(matrix), len(matrix[0])
	mp = make(map[uint64]int)
	minSum := math.MaxInt32

	for x := 0; x < mW; x++ {
		minSum = min(minSum, goUp(matrix, x, mH-1))
	}

	return minSum
}

func printArr(matrix [][]int) {
	fmt.Println()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], ",")
		}
		fmt.Println()
	}
	fmt.Println()
}
func main() {

	arr := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	printArr(arr)
	fmt.Println("minsum", minFallingPathSum(arr))

	arr = [][]int{{-19, 57}, {-40, -5}}
	printArr(arr)
	fmt.Println("minsum", minFallingPathSum(arr))

	arr = [][]int{{-80, -13, 22}, {83, 94, -5}, {73, -48, 61}}
	printArr(arr)
	fmt.Println("minsum", minFallingPathSum(arr))

	arr = [][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}}
	printArr(arr)
	fmt.Println("minsum", minFallingPathSum(arr))
}
