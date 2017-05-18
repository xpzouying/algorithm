package zysort

// sort list by bubble
func bubble_sort(lst []int) []int {

	l := make([]int, len(lst))
	copy(l, lst)

	changed := true

	for i := 0; i < len(l)-1; i++ {
		if !changed {
			break
		}

		for j := 0; j < len(l)-i-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				changed = true
			}
		}

	}

	return l
}
