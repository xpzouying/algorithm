package zysearch

import "container/heap"

/* Find the middle number and return it.
  如果传入的列表有奇数长度，那么返回中间的数值,
  如果传入的列表有偶数长度，那么返回中间两个数的平均值

  因为当为偶数长度的列表时，返回中间两个数的平均值，则利用两个堆来实现。
  一个大堆，一个小堆。

  大顶堆的所有数都小于小顶堆。


  1. 对传入的列表进行入堆处理。
  2. 首先压入大顶堆，压入以后更新大顶堆数据，得到一个新的大顶堆。
  3. 弹出大顶堆的堆顶数据，更新大顶堆内部；
  4. 将弹出的数，压入小顶堆，更新小顶堆内部，得到一个新的小顶堆。
  5. 如果大顶堆的数量大于小顶堆，则弹出小顶堆的堆顶数据并将其压入大顶堆。

  6. 压入所有数据以后，可以得到中间数。
    6.1 如果大顶堆的长度+小顶堆的长度为偶数，则中间数为大顶堆堆顶和小顶堆堆顶的数的平均值；
	6.2 如果大顶堆的长度+小顶堆的长度为奇数，则中间数为大顶堆的堆顶。
*/
func FindTheMiddleNumber(l []int) int {

	// --- init 2 heaps ---
	// min heap
	minHeap := IntMinHeap{}
	heap.Init(&minHeap)

	// max heap
	maxHeap := IntMaxHeap{}
	heap.Init(&maxHeap)

	// Push one element to heaps
	for _, x := range l {
		pushElementToHeaps(&minHeap, &maxHeap, x)
	}

	// Calculate the middle number
	if minHeap.Len() == maxHeap.Len() {
		return (minHeap.Top().(int) + maxHeap.Top().(int)) / 2
	} else {
		return maxHeap.Top().(int)
	}
}

/* 压入一个数压入heaps中

1. push number to max heap
2. pop number from max heap
3. push the pop number to min heap
4. if min heap is longer than max heap,
	then pop one number from min heap,
	and push the number to max heap

5. the middle number
	if the length of two heap is eqaul, then middle number is average of two heap's top number;
	if not equal, then the middle number is max heap's top number.

*/
func pushElementToHeaps(minHeap *IntMinHeap, maxHeap *IntMaxHeap, num int) {
	heap.Push(maxHeap, num)

	popNum := heap.Pop(maxHeap)

	heap.Push(minHeap, popNum)

	// if the length of two heap is not equal
	if maxHeap.Len() < minHeap.Len() {
		transforNum := heap.Pop(minHeap)

		heap.Push(maxHeap, transforNum)
	}

}
