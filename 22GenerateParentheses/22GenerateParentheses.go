package main

import "fmt"

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

func gen(out *[]string, str string, l int, r int) {
	if l == 0 {
		closing := ""
		for i := 0; i < r; i++ {
			closing += ")"
		}
		*out = append(*out, str+closing)
		return
	}

	if l > 0 {
		gen(out, str+"(", l-1, r)
	}
	if l < r {
		gen(out, str+")", l, r-1)
	}
}

func generateParenthesis(n int) []string {
	out := []string{}

	gen(&out, "", n, n)

	return out
}

func main() {
	fmt.Println(generateParenthesis(3))
}
