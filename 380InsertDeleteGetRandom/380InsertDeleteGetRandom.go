package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	occurrences map[int]int
	listOfVals  []int
	size        int
	rnd         *rand.Rand
}

func Constructor() RandomizedSet {
	rs := RandomizedSet{}
	rs.occurrences = make(map[int]int)
	rs.listOfVals = make([]int, 0, 200000)
	rs.size = 0
	rs.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	return rs
}

func (rs *RandomizedSet) Insert(val int) bool {
	if x, ok := rs.occurrences[val]; ok && x != -1 {
		return false
	}
	rs.occurrences[val] = rs.size
	rs.listOfVals = append(rs.listOfVals, val)
	rs.size++
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	i, ok := rs.occurrences[val]
	if !ok || i == -1 { //not exists, or deleted
		return false
	}

	v := rs.listOfVals[rs.size-1] // move last element
	rs.listOfVals[i] = v

	rs.occurrences[v] = i    // update map
	rs.occurrences[val] = -1 // mark deleted
	rs.size--
	rs.listOfVals = rs.listOfVals[:rs.size] // reduce array

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	r := rs.rnd.Intn(rs.size)
	return rs.listOfVals[r]
}

func main() {

	obj := Constructor()
	fmt.Println("remove 0", obj.Remove(0))
	fmt.Println("remove 0", obj.Remove(0))
	fmt.Println("insert 0", obj.Insert(0))
	fmt.Println("random", obj.GetRandom())
	fmt.Println("remove 0", obj.Remove(0))
	fmt.Println("insert 0", obj.Insert(0))
}
