package main

import "fmt"

// Stack implementation using linked list
type stackEl struct {
	val  rune
	prev *stackEl
}

type stack struct {
	top  *stackEl
	size uint
}

func newStack() *stack         { return &stack{(*stackEl)(nil), 0} }
func (s *stack) isEmpty() bool { return s.size == 0 }
func (s *stack) push(c rune) {
	s.top = &stackEl{c, s.top}
	s.size++
}
func (s *stack) pop() rune {
	ret := (s.top).val
	s.top = (s.top).prev
	s.size--
	return ret
}

func isValid(s string) bool {
	st := newStack()
	par := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, c := range s {
		if par[c] != 0 { // if exists in the map
			st.push(c)
		} else {
			if st.isEmpty() { // no opening brackets
				return false
			}
			prev := st.pop()
			if c != par[prev] { // brackets not matchig
				return false
			}
		}
	}
	return st.isEmpty() // check if number of opening and closing match
}

func main() {
	fmt.Println(isValid("("))

}
