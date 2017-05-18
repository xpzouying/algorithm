package zysort

import (
	"fmt"
	"sort"
	"testing"
)

var (
	testSets [][]int
)

func TestBubbleSortNormal(t *testing.T) {
	println(testSets)

	for _, td := range testSets {
		fmt.Printf("orignal list: %v\n", td)
		res := bubble_sort(td)
		fmt.Printf("sorted list: %v\n", res)

		wanted := make([]int, len(td))
		copy(wanted, td)
		sort.Sort(sort.IntSlice(wanted))

		fmt.Printf("Orig: %v ==> %v, zysort result: %v\n", td, wanted, res)

		if !sliceEqual(res, wanted) {
			t.Errorf("Error. %v ==> %v, but result: %v", td, wanted, res)
		}

	}
}

func sliceEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func init() {
	testSets = append(testSets, []int{4, 2, 1, 6})
	testSets = append(testSets, []int{})
	testSets = append(testSets, []int{1})
}
