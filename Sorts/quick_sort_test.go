package zysort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testsets := MockTestset()

	for _, testdata := range testsets {
		QuickSort1(testdata.Data, 0, len(testdata.Data)-1)
		fmt.Println(testdata.Data)

		fmt.Printf("Answer expect: ", testdata.Answer)
	}

}
