package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	K := len(lists)
	if K == 0 {
		return nil
	}
	if K == 1 {
		return lists[0]
	}

	// merge in pairs
	for i := 1; i < K; i++ {
		preHead := &ListNode{}
		prev := preHead
		l1 := lists[0]
		l2 := lists[i]
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				prev.Next = l1
				l1 = l1.Next
			} else {
				prev.Next = l2
				l2 = l2.Next
			}
			prev = prev.Next
		}
		if l1 != nil {
			prev.Next = l1
		} else {
			prev.Next = l2
		}
		lists[0] = preHead.Next
	}
	return lists[0]
}

// classic approach
// func mergeKLists(lists []*ListNode) *ListNode {

// 	K := len(lists)

// 	ptrs := []*ListNode{}
// 	for i := 0; i < K; i++ {
// 		if lists[i] != nil {
// 			ptrs = append(ptrs, lists[i])
// 		}
// 	}

// 	preHead := &ListNode{-1, nil}
// 	prev := preHead

// 	for len(ptrs) > 0 {
// 		minList := -1
// 		minVal := math.MaxInt

// 		for i := 0; i < len(ptrs); i++ {
// 			if ptrs[i].Val <= minVal { //find minimum el
// 				minVal = ptrs[i].Val
// 				minList = i
// 			}
// 		}

// 		if minList != -1 { // if found
// 			ptrs[minList] = ptrs[minList].Next // advance pointer
// 			if ptrs[minList] == nil {          // if reached end of list, remove it
// 				ptrs[minList] = ptrs[len(ptrs)-1]
// 				ptrs = ptrs[:len(ptrs)-1]
// 			}
// 			prev.Next = &ListNode{minVal, nil} // add el to output list
// 			prev = prev.Next
// 		}
// 	}
// 	return preHead.Next
// }

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
