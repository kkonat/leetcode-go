package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func minimumAverageDifference(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	lsum, rsum := nums[0], 0

	for i := 1; i < len(nums); i++ {
		rsum += nums[i]
	}

	lnum, rnum := 1, len(nums)-1
	lavg, ravg := lsum, rsum/rnum

	minavgI := 0
	minavg := abs(lavg - ravg)

	for i := 1; i <= len(nums)-1; i++ {
		lnum++
		rnum--
		lsum += nums[i]
		rsum -= nums[i]
		lavg = lsum / lnum
		if rnum > 0 {
			ravg = rsum / rnum
		} else {
			ravg = 0
		}
		t := abs(lavg - ravg)
		if t < minavg {
			minavg = t
			minavgI = i
		} else if t == minavg {
			if i < minavgI {
				minavgI = i
			}
		}
	}
	return minavgI
}

func main() {
	fmt.Println(minimumAverageDifference([]int{4, 2, 0}))
}
