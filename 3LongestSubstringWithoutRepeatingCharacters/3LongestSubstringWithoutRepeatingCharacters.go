package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	ls := len(s)
	max := 0

	for l := 0; l < ls; l++ { // try build substrings sarting from the left side
		nonUnique := make([]bool, 256)
		for j := l; j < ls; j++ { // try l characters
			r := int(s[j])
			if !nonUnique[r] {
				nonUnique[r] = true
			} else {
				t := j - l
				if t > max {
					max = t
				}
				break
			}
			t := j - l + 1
			if t > max {
				max = t
			}
		}
	}
	return max
}
func main() {
	fmt.Println("longest substring: ", lengthOfLongestSubstring("uwawurthuwawurthuwawurth"))
}
