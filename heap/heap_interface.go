package heap

import (
	"container/heap"
	"fmt"
)

// 通过实现接口的方式 小根/大根堆
// 自定义类型 实现heap.Interface 接口
// 然后使用heap.Init(h),heap.Push(h, 3),poppedValue := heap.Pop(h)等方法
// 堆排序时间复杂度：
// 建堆heap.Init： O(n)
// 			最后一个非叶子节点（位于索引 n/2 - 1 处）开始，反复进行 "向下冒泡" 操作）
// 堆排序： O(nlogn)
//			每次从堆中取出根节点，将根节点与堆的最后一个节点交换，然后调整堆，使其重新满足堆的性质。这个过程需要 log n 次比较和交换，而堆中有 n 个元素，所以总体的时间复杂度是 O(n log n)
// Push 操作： O(logn)
//			插入元素，将元素添加到堆的最后，然后执行 "向上冒泡" 操作，最坏情况时间复杂度是堆的高度 O(log n)。
// Pop 操作： O(logn)
//			删除堆顶元素后，将堆的最后一个元素移到堆顶，然后执行 "向下冒泡" 操作

type heapInt []int

func (h *heapInt) Len() int {
	return len(*h)
}
func (h *heapInt) Less(i, j int) bool {
	// 小根堆
	return (*h)[i] < (*h)[j]
	// 大根堆
	// return (*h)[i] > (*h)[j]
}
func (h *heapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
func InterfaceExample() {
	// 使用指针传递 覆盖原始值
	h := &heapInt{5, 6, 7, 4, 5, 6, 7, 9, 12}
	heap.Init(h)
	heap.Push(h, 1)
	fmt.Println(heap.Pop(h)) // 1
	fmt.Println(heap.Pop(h)) // 4
}
