package zysearch

import "testing"

// Test Find the middle number with normal slice
func TestFindMiddleNumberNormalList1(t *testing.T) {
	l := []int{8, 1, 3, 2, 5, 9}
	expectRes := 4

	m := FindTheMiddleNumber(l)

	if m != expectRes {
		t.Errorf("Result expect be: %d, but got result be: %d", expectRes, m)
	}
}

// Test Find the middle number with normal slice
func TestFindMiddleNumberNormalList2(t *testing.T) {
	l := []int{8, 1, 3, 2, 5, 9, 6}
	expectRes := 5

	m := FindTheMiddleNumber(l)

	if m != expectRes {
		t.Errorf("Result expect be: %d, but got result be: %d", expectRes, m)
	}
}

// Test Find the middle number with normal slice
func TestFindMiddleNumberNormalListOnlyOneElement(t *testing.T) {
	l := []int{8}
	expectRes := 8

	m := FindTheMiddleNumber(l)

	if m != expectRes {
		t.Errorf("Result expect be: %d, but got result be: %d", expectRes, m)
	}
}
