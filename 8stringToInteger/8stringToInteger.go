package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	start := 0
	neg := false
	for start < len(s) && s[start] == ' ' {
		start++
	}
	if start >= len(s) {
		return 0
	}

	if s[start] == '+' {
		start++
		goto cont
	}
	if start >= len(s) {
		return 0
	}

	if s[start] == '-' {
		start++
		neg = true
	}
cont:
	end := start
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	if start == end {
		return 0
	}

	var limit int64
	if neg {
		limit = -math.MinInt32
	} else {
		limit = math.MaxInt32
	}
	var value int64

	for i := start; i < end; i++ {
		value *= 10
		value += int64(s[i]) - int64('0')
		if value > limit {
			value = limit
			break
		}
	}
	if neg {
		value = -value
	}
	return int(value)
}

func main() {
	fmt.Println(myAtoi(""))
}
