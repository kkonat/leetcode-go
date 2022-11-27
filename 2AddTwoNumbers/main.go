package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func llen(head *ListNode) int {

	// calc length
	ptr := head
	var llen int = 0
	for ptr != nil {
		llen++
		ptr = ptr.Next
	}
	return llen
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	p1, p2 := l1, l2

	var max int
	ll1, ll2 := llen(l1), llen(l2)
	if ll1 > ll2 {
		max = ll1
	} else {
		max = ll2
	}

	var head *ListNode = nil
	var prev *ListNode = nil
	var node *ListNode = nil

	carry := 0
	sum := 0
	res := 0

	for i := 0; i < max+1; i++ {
		sum = carry
		if p1 != nil {
			sum += p1.Val
		}
		if p2 != nil {
			sum += p2.Val
		}
		res = sum % 10
		carry = (sum - res) / 10
		if p1 != nil || p2 != nil || res != 0 {
			node = &ListNode{res, nil}
		} else {
			node = nil
		}
		if i == 0 {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	return head
}

func main() {
	fmt.Println("2AddTwoNumbers")
	// l1 := createList([]int{9, 9, 9, 9, 9, 9, 9})
	// l2 := createList([]int{9, 9, 9, 9})
	// expected: [8,9,9,9,0,0,0,1]
	//l1 := createList([]int{2, 4, 9})
	//l2 := createList([]int{5, 6, 4, 9})
	// expected: [7,0,4,0,1]
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	// expected: [7,0,8]
	fmt.Println("l1:", l1.toString())
	fmt.Println("l2:", l2.toString())

	result := addTwoNumbers(l1, l2)

	fmt.Println("result:", result.toString())
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

func createList(a []int) *ListNode {
	head := &ListNode{a[0], nil}
	prev := head
	for i := 1; i < len(a); i++ {
		prev.Next = &ListNode{a[i], nil}
		prev = prev.Next
	}
	return head
}
