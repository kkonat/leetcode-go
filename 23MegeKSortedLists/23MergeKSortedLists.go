package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	K := len(lists)

	ptrs := make([]*ListNode, K)
	for i := 0; i < K; i++ {
		ptrs[i] = lists[i]
	}

	var head *ListNode = nil
	var prev *ListNode = nil

	minList := -1
	minval := math.MaxInt

	TotalNilPtrs := 0

	for TotalNilPtrs < K {
		for i := 0; i < K; i++ {
			TotalNilPtrs = 0
			if ptrs[i] == nil {
				TotalNilPtrs++
			} else {
				if ptrs[i].Val < minval { //find minimum el
					minList = i
					minval = ptrs[i].Val
				}
			}
		}
		fmt.Println("minList: ", minList, "minval=", minval)
		n := &ListNode{minval, nil}        //create new Node
		ptrs[minList] = ptrs[minList].Next // advance pointer
		if prev != nil {
			prev.Next = n
		} else {
			head = n
		}
		prev = n
	}

	return head
}

func createLists(lists [][]int) []*ListNode {

	listsArr := make([]*ListNode, len(lists))

	for i := 0; i < len(lists); i++ {
		var prev *ListNode = nil
		for j := 0; j < len(lists[i]); j++ {
			n := &ListNode{lists[i][j], nil}
			if j != 0 {
				prev.Next = n
			} else {
				listsArr[i] = n
			}
			prev = n
		}
	}
	return listsArr
}
func (list *ListNode) toString() string {
	str := "["
	for ptr := list; ptr != nil; ptr = ptr.Next {
		str += " " + fmt.Sprintf("%d", ptr.Val)
	}
	str += "]"
	return str

}
func main() {
	lists := createLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	// for i := 0; i < len(lists); i++ {
	// 	fmt.Printf("List%d: %s\n", i, lists[i].toString())
	// }`

	result := mergeKLists(lists)
	fmt.Println("result: ", result.toString())

}
