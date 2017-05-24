# 查找算法 #

包含各类查找相关的算法

## List ##

- [ ] 查找中位数，如果列表长度为偶数，则取len(list)/2。
- [ ] 查找中位数，如果列表长度为偶数，则取中间两个数的平均值。
- [x] 查找第K小/大的数。


## 技巧 ##


### 查找中位数，如果列表长度为偶数，则取len(list)/2 ###

该问题不需要求平均值，那么可以使用`查找第K小/大的数`计算。K为len(list)/2。



### 查找中位数，如果列表长度为偶数，则取中间两个数的平均值。 ###

该问题需要求中间两个数的平均值，那么使用最小堆和最大堆。

缺点：
- 使用堆的问题是，内存需要足够放下所有的数据。因为需要将所有的数都存放在堆中，堆需要保存在内存中。



### 查找第K小/大的数 ###

以查找第K小的数为例，

1. 类似于快速排序的第一步，挑选数组中的一个数X，将数组进行partition操作，X前面的数都小于X，X后面的数都大于X值。

排序完后为 [1, 2, 3, ... X ... N-1, N]，

2. 当X的index大于K，那么表示第K小的数还在X的前面，那么我们就舍弃掉X以及X后面的数，在列表的前面部分查找第K小的值，

3. 当X的index小于K，那么表示第K小的数在X的后面，那么我们就舍弃掉X以及X前面的数，我们在列表的后面部分查找第K小的值。

4. 递归计算，直到找到第K小的值为止。


代码参考：
- the_smallest_k_number.go
- the_smallest_k_number_test.go