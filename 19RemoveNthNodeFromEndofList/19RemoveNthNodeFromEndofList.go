package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	N := n + 1
	ptrs := make([]*ListNode, N)
	in := 0

	for p := head; p != nil; p = p.Next {
		ptrs[in] = p
		in = (in + 1) % N
	}
	p := ptrs[(in)%N]
	fmt.Println("p:", p.Val)
	t := ptrs[(in)%N].Next
	fmt.Println("t:", t.Val)

	p.Next = t.Next
	return head
}

func main() {
	fmt.Println(removeNthFromEnd(createList([]int{1, 2, 3, 4, 5}), 2).toString())
	fmt.Println(removeNthFromEnd(createList([]int{1, 2}), 1).toString())
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

func (list *ListNode) toString() string {
	str := "["
	for ptr := list; ptr != nil; ptr = ptr.Next {
		str += " " + fmt.Sprintf("%d", ptr.Val)
	}
	str += "]"
	return str

}
