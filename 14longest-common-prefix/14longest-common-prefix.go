package main

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if strs[0] == "" {
		return ""
	}
	for max := 0; max <= 200; max++ { // iterate throug max 201 chars
		if max >= len(strs[0]) {
			return strs[0]
		}
		for n := 1; n < len(strs); n++ { // see if it's the same in all strings
			// if one of the strings is shorter OR one of the chars differs, end
			if len(strs[n]) < max+1 || rune(strs[n][max]) != rune(strs[0][max]) {
				if max == 0 {
					return ""
				} else {
					return strs[0][0:max]
				}
			}
		}
	}
	return ""
}

// @lc code=end
