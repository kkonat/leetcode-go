package main

import (
	"encoding/json"
	"fmt"
)

func scanAndBuld(scan *TreeNode, build *TreeNode) {

	if scan.Left != nil {
		build.Left = &TreeNode{0, nil, nil}
		scanAndBuld(scan.Left, build.Left)
		build.Val += build.Left.Val
	}
	if scan.Right != nil {
		build.Right = &TreeNode{0, nil, nil}
		scanAndBuld(scan.Right, build.Right)
		build.Val += build.Right.Val
	}
	build.Val += scan.Val
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func findMax(maxProd *int, total int, scan *TreeNode) {
	if scan.Left != nil {
		fmt.Println("l:", scan.Left.Val, (total - scan.Left.Val))
		*maxProd = max(*maxProd, scan.Left.Val*(total-scan.Left.Val))
		findMax(maxProd, total, scan.Left)
	}
	if scan.Right != nil {
		fmt.Println("r:", scan.Right.Val, (total - scan.Right.Val))

		*maxProd = max(*maxProd, scan.Right.Val*(total-scan.Right.Val))
		findMax(maxProd, total, scan.Right)
	}
}

func maxProduct(root *TreeNode) int {
	// 1. scan tree, calc total sum and subsums in each node
	subsumsRoot := &TreeNode{0, nil, nil}
	scanAndBuld(root, subsumsRoot)
	fmt.Println("Total sum:", subsumsRoot.Val)
	PrintTree("2", subsumsRoot)
	maxVal := 0
	// 2. scan subsums tree and find max of (total-subsum)*subsum
	findMax(&maxVal, subsumsRoot.Val, subsumsRoot)

	return maxVal
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromJsonString(inputStr string) *TreeNode {
	var input []*int
	if err := json.Unmarshal([]byte(inputStr), &input); err != nil {
		panic(err)
	}
	if len(input) == 0 {
		return nil
	} else if len(input) == 1 {
		return &TreeNode{*input[0], nil, nil}
	}
	r := &TreeNode{*input[0], nil, nil}
	input = input[1:]
	stack := []*TreeNode{r}
	stack1 := []*TreeNode{}
	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		if len(input) != 0 {
			i := input[0]
			input = input[1:]
			if i != nil {
				p.Left = &TreeNode{*i, nil, nil}
				stack1 = append(stack1, p.Left)
			}
		}
		if len(input) != 0 {
			i := input[0]
			input = input[1:]
			if i != nil {
				p.Right = &TreeNode{*i, nil, nil}
				stack1 = append(stack1, p.Right)
			}
		}
		if len(stack) == 0 {
			stack = stack1
			stack1 = []*TreeNode{}
		}
	}
	return r
}

func PrintTree(prefix string, n *TreeNode) {
	if n != nil {
		PrintTree(prefix+"    ", n.Right)
		fmt.Printf("%s #-- %d\n", prefix, n.Val)
		PrintTree(prefix+"    ", n.Left)
	}
}

func main() {
	// root := CreateTreeFromJsonString("[1,null,2,3,4,null,null,5,6]")
	root := CreateTreeFromJsonString("[1,2,3,4,5,6]")
	PrintTree("1", root)
	fmt.Println("MaxProduct:", maxProduct(root))
}
