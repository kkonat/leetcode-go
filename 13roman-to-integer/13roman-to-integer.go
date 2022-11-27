package main

// /*
//  * @lc app=leetcode id=13 lang=golang
//  *
//  * [13] Roman to Integer
//  */

// @lc code=start

func romanToInt(s string) int {
	vals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	chars := []rune(s)

	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		val := vals[chars[i]]
		if i > 0 {
			prev := vals[chars[i-1]]

			if (val == 5 || val == 10) && prev == 1 {
				val -= 1
				i--
			}
			if (val == 50 || val == 100) && prev == 10 {
				val -= 10
				i--
			}
			if (val == 500 || val == 1000) && prev == 100 {
				val -= 100
				i--
			}
		}
		result += val
	}
	return result
}

// @lc code=end
