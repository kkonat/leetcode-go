/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func minR(arr []int, i int) int {
	if i < len(arr) {
		return arr[i]
	} else {
		return math.MaxInt
	}
}
func maxL(arr []int, i int) int {
	if i > 0 {
		return arr[i-1]
	} else {
		return math.MinInt
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	arr1Len := len(nums1)
	arr2Len := len(nums2)
	noElms := arr1Len + arr2Len

	start := 0
	end := arr1Len
	for start <= end {
		p1 := start + (end-start)/2
		p2 := (noElms+1)/2 - p1

		l1 := maxL(nums1, p1)
		l2 := maxL(nums2, p2)
		r1 := minR(nums1, p1)
		r2 := minR(nums2, p2)

		if l1 <= r2 && l2 <= r1 {
			med1 := 0
			med2 := 0
			if l1 > l2 {
				med1 = l1
			} else {
				med1 = l2
			}
			if noElms%2 != 0 {
				return float64(med1)
			} else {
				if r1 < r2 {
					med2 = r1
				} else {
					med2 = r2
				}
				return float64(med1+med2) / 2.0
			}
		}
		if l1 > r2 {
			end = p1 - 1
		} else {
			start = p1 + 1
		}
	}
	return 0
}

// @lc code=end

func main() {
	fmt.Println("Median = ", findMedianSortedArrays([]int{1, 3}, []int{2}), "Expected:", 2.0)
	fmt.Println("Median = ", findMedianSortedArrays([]int{1, 2}, []int{3, 4}), "Expected:", 2.5)
	fmt.Println("Median = ", findMedianSortedArrays([]int{5, 7, 9, 10}, []int{4, 5, 6, 9}), "Expected:", 6.5)
	fmt.Println("Median = ", findMedianSortedArrays([]int{3, 7, 10}, []int{4, 5, 6, 9}), "Expected:", 6)

}
