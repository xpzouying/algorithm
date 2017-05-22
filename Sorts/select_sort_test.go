package zysort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	testsets := MockTestset()

	for _, testdata := range testsets {
		data := make([]int, len(testdata.Data))
		copy(data, testdata.Data)

		SelectSort(testdata.Data)
		fmt.Printf("Orignal: %v, sorted by zouying: %v, answer: %v\n", data, testdata.Data, testdata.Answer)

		fmt.Printf("Orignal addr : %p, new data addr: %p,\n", &testdata.Data, &data)

		if !sliceEqual(testdata.Data, testdata.Answer) {
			t.Errorf("Orignal: %v, sorted by zouying: %v, answer: %v\n", data, testdata.Data, testdata.Answer)
		}
	}

}
