package zysearch

import "fmt"

/*Find the biggest k number in slice

Select the first number to pivot, which to split the slice to two part.
The left is smaller than pivot, the right part is greater than pivot.

Returns:
	If not found, then return -1.
	Else return the right position, start from 0.
*/
func FindSmallestKNumber(K int, L []int) int {

	// 查找错误
	if len(L) <= 0 || K > len(L) {
		fmt.Printf("List: %v, K=%d\n", L, K)

		return -1
	}

	pivot := partition(L, 0, len(L)-1)

	fmt.Println("After partition. Pivot=", pivot, " slice= ", L, " K=", K)

	if pivot == K-1 {
		fmt.Printf("Find the small K=%d, value=%d\n", K, L[pivot])
		return L[pivot]
	} else if pivot > K-1 {
		// 如果找到的是 [0,1,2,...,K,K+1,...(pivot)...N]
		// 那么在左边部分继续找K

		// 右侧部分全部舍弃，舍弃部分的数量为：pivot值+右侧部分所有的值
		// 新的列表长度为：0, 1, 2,...,pivot
		slc := make([]int, pivot)
		copy(slc, L[:pivot])
		fmt.Println("New slice: ", slc)
		return FindSmallestKNumber(K, slc)
	} else {
		// 如果找到的是 [0,1,2,...(pivot)...,K,K+1,...N]
		// 那么在右侧部分继续找K

		// 左侧部分全部舍弃，舍弃部分数量为：pivot数值加上左侧部分，
		// 为：1+pivot
		leftPart := 1 + pivot
		rightPart := len(L) - leftPart
		slc := make([]int, rightPart)
		copy(slc, L[pivot+1:])
		fmt.Println("New slice: ", slc)

		// 需要在右侧找到第N大的值，N=K-leftPart
		return FindSmallestKNumber(K-leftPart, slc)
	}

	return 0
}

/*
Partition the slice/list.

Select the first one to be pivot.

	Returns:
		Return -1: not found the right position.
		Return other index: right index.
*/
func partition(L []int, lo, hi int) int {

	// select the last one to be pivot
	pivotV := L[hi]

	// the left is all smaller than pivot
	wall := lo

	// 为了更好的查看逻辑，没有将index写在for循环中
	for index := lo; index < hi; {

		if pivotV > L[index] {
			L[index], L[wall] = L[wall], L[index]
			wall++
			index++
		} else if pivotV <= L[index] {
			index++
		}
	}

	L[wall], L[hi] = L[hi], L[wall]

	return wall
}
