package zysort

type TestData struct {
	Data   []int
	Answer []int
}

var (
	emptydata  = TestData{[]int{}, []int{}}
	onedata    = TestData{[]int{0}, []int{0}}
	twodata    = TestData{[]int{0, 2}, []int{0, 2}}
	tworevdata = TestData{[]int{3, 2}, []int{2, 3}}
	orderdata  = TestData{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}
	randdata   = TestData{[]int{3, 1, 2, 5, 4, 8, 6}, []int{1, 2, 3, 4, 5, 6, 8}}
)

func MockTestset() (testdata []TestData) {
	testdata = append(testdata, emptydata, onedata, twodata, tworevdata, orderdata, randdata)

	return testdata
}
