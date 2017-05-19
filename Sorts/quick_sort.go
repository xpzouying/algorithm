package zysort

import "fmt"

/*快速排序


 */

/* QuickSort1: 使用迭代的方法进行快速排序

Args:
	l: 需要进行排序的列表
	b: 开始索引
	e: 结束索引

Notes:
	在原有列表的基础上进行排序。

	Tips:
	快速排序为分治法，所以技巧在于如何分治。
	主要是找到划分两块区域的界线，然后分别对两个小区域进行快速排序。

	<1> div := partition(list, begin, end)
	<2> quicksort(list, begin, div)
	<3> quicksort(list, div, end)
*/
func QuickSort1(l []int, b, e int) {

	fmt.Printf("Begin: %d, end: %d, list: %v\n", b, e, l)

	if len(l) <= 1 {
		return
	}

	if b >= e {
		return
	}

	// fmt.Println("Origin list: ", l)
	div := partition(l, b, e)
	// fmt.Println("After partition, list: ", l)
	if div > 0 {
		QuickSort1(l, div-1, e)
	}

	if div < e {
		QuickSort1(l, b, div-1)
	}
}

/*Partition list.

Choose the first one element, povit.
Split the list, move the smaller (than povit element) to left,
move the bigger element to right.
*/
func partition(l []int, b, e int) (d int) {
	// use the first element to be povit
	povit := l[b]
	i, j := b+1, e

	for i <= j {
		// find the small element in left
		for povit <= l[j] && i <= j {
			j--
		}

		for l[i] < povit && i <= j {
			i++
		}

		if i < j {
			l[i], l[j] = l[j], l[i]
		}
	}

	l[i-1], l[b] = l[b], l[i-1]
	return i - 1
}
