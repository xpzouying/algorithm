package zysort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testsets := MockTestset()

	for _, testdata := range testsets {
		data := make([]int, len(testdata.Data))
		copy(data, testdata.Data)

		quicksort_by_wall(testdata.Data, 0, len(testdata.Data)-1)
		fmt.Printf("Orignal: %v, sorted by zouying: %v, answer: %v\n", data, testdata.Data, testdata.Answer)

		if !sliceEqual(testdata.Data, testdata.Answer) {
			t.Errorf("Orignal: %v, sorted by zouying: %v, answer: %v\n", data, testdata.Data, testdata.Answer)
		}
	}

}
