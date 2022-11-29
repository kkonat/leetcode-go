package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(e1 **ListNode) {
	olde1 := (*e1)
	e2 := olde1.Next
	e3 := e2.Next

	(*e1).Next = e3
	e2.Next = olde1
	*e1 = e3
}
func swapPairs(head *ListNode) *ListNode {

	preHead := &ListNode{0, head}

	ptr := preHead.Next

	for ptr != nil && ptr.Next != nil {
		swap(&ptr)
	}
	return preHead.Next
}

func createList(a []int) *ListNode {
	head := &ListNode{a[0], nil}
	prev := head
	for i := 1; i < len(a); i++ {
		prev.Next = &ListNode{a[i], nil}
		prev = prev.Next
	}
	return head
}
func (l *ListNode) toString() string {
	str := "["
	ptr := l
	for ptr != nil {
		str += fmt.Sprint(ptr.Val) + " "
		ptr = ptr.Next
	}
	str += "]"
	return str
}

func main() {
	h := createList([]int{1, 2, 3, 4})
	fmt.Println(h.toString())
	swapPairs(h)
	fmt.Println(h.toString())
}
