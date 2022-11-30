package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	oc := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		oc[arr[i]]++
	}
	un := make(map[int]bool)
	for _, val := range oc {
		if un[val] {
			return false
		} else {
			un[val] = true
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2}))
}
