package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start

func rev(x int) string {
	strx := fmt.Sprintf("%d", x)
	rev := ""
	for i := 0; i < len(strx); i++ {
		rev += string(strx[len(strx)-1-i])
	}
	return rev
}

func conv(strx string, neg bool) int {
	result := 0
	mul := 1
	l := len(strx)
	for i := 0; i < l; i++ {
		result += int(strx[l-1-i]-'0') * mul
		mul *= 10
	}
	if neg {
		return -result
	} else {
		return result
	}
}

func lessThan(x, limit string) bool {
	if len(x) < len(limit) {
		return true
	} else if len(x) > len(limit) {
		return false
	} else {
		for i := 0; i < len(x); i++ {
			if x[i] < limit[i] {
				return true
			} else if x[i] > limit[i] {
				return false
			}
		}
		return false
	}
}
func reverse(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}

	strx := rev(x)
	min := fmt.Sprintf("%d", -math.MinInt32)
	max := fmt.Sprintf("%d", math.MaxInt32)

	if neg {
		if lessThan(strx, min) {
			return conv(strx, neg)
		} else {
			return 0
		}
	} else {
		if lessThan(strx, max) {
			return conv(strx, neg)
		} else {
			return 0
		}
	}
}

// @lc code=end

func main() {
	fmt.Println("result=", reverse(1534236469))  // expected 0
	fmt.Println("result=", reverse(1563847412))  // expected 0
	fmt.Println("result=", reverse(-1563847412)) // expected 0
	fmt.Println("result=", reverse(-8463847421)) // expect ok
	fmt.Println("result=", reverse(123))         // expect ok

}
