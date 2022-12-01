package main

import "fmt"

func halvesAreAlike(s string) bool {
	l := len(s) / 2
	c1 := 0
	c2 := 0
	v := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for i := 0; i < l; i++ {
		c := s[i]
		if v[c] || v[c+byte(32)] {
			c1++
		}
		c = s[l+i]
		if v[c] || v[c+byte(32)] {
			c2++
		}
	}
	return c1 == c2
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
