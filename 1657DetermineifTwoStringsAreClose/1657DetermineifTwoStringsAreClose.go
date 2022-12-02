package main

import "fmt"

func closeStrings(word1 string, word2 string) bool {
	l1, l2 := len(word1), len(word2)
	if l1 != l2 {
		return false
	}
	if l1 == 0 {
		return false
	}

	c1 := map[byte]int{}
	c2 := map[byte]int{}
	for i := 0; i < l1; i++ { // count how often each letter occurs
		c1[word1[i]]++ // e.g for cdeaaabbdddd  you get: c:1, d:1, e:1, a:3, b2:2 d:4
		c2[word2[i]]++
	}

	//check if both words contain the same letters
	for k, _ := range c1 { //take each letter in c1
		if _, ok := c2[k]; !ok { // if it doesn't exist in c2, fail
			return false
		}
	}

	h := map[int]int{}
	total := 0
	for _, v := range c2 { // count how many letter sequences of various leghts there ara
		h[v]++  // e.g for the above example, you get 1:3, 2:1, 3:1, 4:1
		total++ // incr. no. of total sequences of various leghts
	}

	for _, v := range c1 { // take the other word and see if the number of sequences is the same
		if h[v] > 0 {
			h[v]--
			total--
		} else {
			return false
		}
	}
	return total == 0 // if the other word had different numver of sequences
}

func main() {
	fmt.Println(closeStrings("uau", "ssx"))
}
