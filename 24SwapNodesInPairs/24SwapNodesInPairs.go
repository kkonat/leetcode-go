package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(prev **ListNode) {
	e1 := (*prev).Next
	e2 := e1.Next
	if e2 == nil {
		return
	}
	e3 := e2.Next

	(*prev).Next = e2
	e2.Next = e1
	e1.Next = e3
}

func swapPairs(head *ListNode) *ListNode {

	preHead := &ListNode{0, head}

	for ptr := preHead; ptr != nil && ptr.Next != nil; ptr = ptr.Next {
		swap(&ptr)
		ptr = ptr.Next
		if ptr.Next == nil {
			return preHead.Next
		}
	}
	return preHead.Next
}

func createList(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
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
	if l != nil {
		ptr := l
		for ptr != nil {
			str += fmt.Sprint(ptr.Val) + " "
			ptr = ptr.Next
		}
	}
	str += "]"
	return str
}

func main() {
	h := createList([]int{1, 2, 3, 4})
	fmt.Println(h.toString())
	h = swapPairs(h)
	fmt.Println(h.toString())
}
