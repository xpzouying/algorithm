package zysort

// select sort
func SelectSort(lst []int) {

	for i := 0; i < len(lst)-1; i++ {

		// pre-select the first to smallest
		min := i

		j := i

		for ; j < len(lst); j++ {
			if lst[min] > lst[j] {
				min = j
			}
		}

		if min != i {
			// swap it
			lst[min], lst[j] = lst[j], lst[min]
		}
	}
}
