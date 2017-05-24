package zysearch

import (
	"testing"
)

func TestFindBiggestKNumber1(t *testing.T) {
	L := []int{5, 2, 1, 8, 6, 3}

	res := FindSmallestKNumber(3, L)

	if res != 3 {
		t.Errorf("K=3, number should be *3*, the result=%d\n", res)
	}

	t.Log(res)
}

func TestFindSmallestKNumberInLeft(t *testing.T) {
	L := []int{5, 2, 1, 8, 6, 3}

	res := FindSmallestKNumber(2, L)

	if res != 2 {
		t.Errorf("K=2, number should be *2*, the result=%d\n", res)
	}

	t.Log(res)
}

func TestFindSmallestKNumberInRight(t *testing.T) {
	L := []int{5, 2, 1, 8, 6, 3}

	res := FindSmallestKNumber(4, L)

	if res != 5 {
		t.Errorf("K=4, number should be *4*, the result=%d\n", res)
	}

	t.Log(res)
}

func TestFindBiggestKNumberOnlyOneNumber(t *testing.T) {
	L := []int{5}

	res := FindSmallestKNumber(1, L)
	if res != 5 {
		t.Errorf("K=1, number should be *5*, the result=%d\n", res)
	}

	t.Log(res)
}
