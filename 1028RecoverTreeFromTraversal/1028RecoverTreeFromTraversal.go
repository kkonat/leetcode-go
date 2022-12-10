package main

import (
	"bytes"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func parseVal(p *string, i int) int32 {

}
func parseLvl(p string, i int) int32 {

}

func recoverFromPreorder(traversal string) *TreeNode {
	var buf bytes.Buffer

	i := 0
	l := len(traversal)
	for i < l {
		for ; p[i] >= '0' && p[i] < '9' && i < l; i++ {
			buf.WriteByte(p[i])
		}
		if v, err := strconv.ParseInt(buf.String(), 10, 32); err == nil {
			return int32(v)
		} else {
			panic(err)
		}
		lv := 0
		for ; p[i] == '-'; i++ {
			lv++
		}
	}
	return lv
}
