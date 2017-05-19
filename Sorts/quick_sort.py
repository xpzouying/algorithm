# -*- coding: utf-8 -*-


''' Way 2: quick sort by wall

    Partition the list by wall.
    Ref:
        http://www.99nth.com/~krm/blog/quicksort-python.html
'''

def partition_by_wall(L, lo, hi):

    '''
    对于有墙的快速排序，技巧在于：
    wall假设最开始在最左边，移动坐标从左开始搜索到列表的最后，
    移动坐标所在的值如果小于pivot的话，就表示正常，wall跟着移动；
    当发现移动坐标大于pivot的话，墙体不移动，游动坐标继续前进，直到找到一个小于pivot的值，与wall进行交换。
    '''

    # 使用最后一个值作为pivot
    pivot = L[hi]

    # index墙，低于墙的值 都小于pivot
    wall = lo

    # index： 移动/游动坐标
    index = lo

    while index < hi:

        if L[index] <= pivot:
            L[wall], L[index] = L[index], L[wall]

            wall += 1

        index += 1

    # 循环结束后，wall坐标确定，wall左边的都比pivot小，wall右边的都比pivot大，
    # swap wall <--> hi
    L[wall], L[hi] = L[hi], L[wall]
    return wall


def quicksort_by_wall(L, lo, hi):
    if len(L) <= 1:
        return L

    if lo >= hi:
        return L

    pivot = partition_by_wall(L, lo, hi)
    quicksort_by_wall(L, lo, pivot-1)
    quicksort_by_wall(L, pivot+1, hi)
