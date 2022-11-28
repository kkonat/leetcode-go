/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package main

import (
	"fmt"
)

// @lc code=start
func minR(arr []int, i int) int {
	if i == len(arr) {
		return -1001
	} else {
		return arr[i]
	}
}
func maxL(arr []int, i int) int {
	if i == len(arr) {
		return 1001
	} else {
		return arr[i]
	}
}
func medOf(arr []int) float64 {
	l := len(arr)
	if l%2 == 0 {
		return float64(arr[l/2-1]+arr[l/2]) / 2.0
	} else {
		return float64(arr[l/2])
	}
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Printf("nums1: %v, nums2: %v\n", nums1, nums2)

	if len(nums1) == 0 && len(nums2) == 0 {
		return -1
	}

	if len(nums1) == 0 {
		return medOf(nums2)
	}

	if len(nums2) == 0 {
		return medOf(nums1)
	}

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	a1l := len(nums1)
	a2l := len(nums2)
	k := a1l + a2l
	ll := (k + 1) / 2

	low := 0
	hai := a1l
	for low <= hai {
		nums1_count := low + (hai-low)/2
		nums2_count := ll - nums1_count

		maxLeft1 := maxL(nums1, nums1_count)
		maxLeft2 := maxL(nums2, nums2_count)
		minRight1 := minR(nums1, nums1_count)
		minRight2 := minR(nums2, nums2_count)

		fmt.Println("[", maxLeft1, minRight1, "] [", maxLeft2, minRight2, "]")
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			med1 := 0
			med2 := 0
			if maxLeft1 > maxLeft2 {
				med1 = maxLeft1
			} else {
				med1 = maxLeft2
			}
			if k%2 != 0 {
				return float64(med1)
			} else {
				if minRight1 < minRight2 {
					med2 = minRight1
				} else {
					med2 = minRight2
				}
				return float64(med1+med2) / 2.0
			}
		}
		if maxLeft1 > minRight2 {
			hai = nums1_count - 1
		} else {
			low = nums1_count + 1
		}
	}
	return -2
}

// @lc code=end

func main() {
	// fmt.Println("Median = ", findMedianSortedArrays([]int{1, 3}, []int{2}), "Expected:", 2.0)
	// fmt.Println("Median = ", findMedianSortedArrays([]int{1, 2}, []int{3, 4}), "Expected:", 2.5)
	fmt.Println("Median = ", findMedianSortedArrays([]int{5, 7, 9, 10}, []int{4, 5, 6, 9}), "Expected:", 6.5)
	// fmt.Println("Median = ", findMedianSortedArrays([]int{3, 7, 10}, []int{4, 5, 6, 9}), "Expected:", 6)

}
