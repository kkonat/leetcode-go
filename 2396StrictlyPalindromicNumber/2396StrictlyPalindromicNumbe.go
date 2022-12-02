package main

//An integer n is strictly palindromic if, for every base b between 2 and n - 2 (inclusive),
// the string representation of the integer n in base b is palindromic.
// Constraints:

// 4 <= n <= 10^5
func ispalindromic(v []int) bool {
	l := len(v)
	for i := 0; i < l/2; i++ {
		if v[i] != v[l-1-i] {
			return false
		}
	}
	return true
}
func reBase(n int, b int) []int {
	rb := []int{}
	if n == 0 {
		return []int{0}
	}
	for n > 0 {
		d := n % b
		rb = append(rb, d) // the array is reversed, but here it doesn't matter
		n /= b
	}
	return rb
}
func isStrictlyPalindromic(n int) bool {
	for b := 2; b <= n-2; b++ {
		rb := reBase(n, b)
		if !ispalindromic(rb) {
			return false
		}
	}

	return true
}
