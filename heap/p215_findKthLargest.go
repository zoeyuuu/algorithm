package heap

import (
	"container/heap"
	"fmt"
)

// 215. 数组中的第K个最大元素 2023-12-07 422 !!!
// https://leetcode.cn/problems/kth-largest-element-in-an-array/?envType=study-plan-v2&envId=top-100-liked
// 方法一：自己构建大顶堆 面试中经常考 自己实现 findKthLargest1 (本题只用到BuildMaxHeap、pop、down三个函数)
// 方法二：实现接口heap.Interface findKthLargest2 没特殊要求实现接口简单一些

func Problem215() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 4
	fmt.Println(findKthLargest1(nums, k))
	fmt.Println(findKthLargest2(nums, k))
}
func findKthLargest1(nums []int, k int) int {
	BuildMaxHeap(nums)
	fmt.Println(nums)
	for i := 1; i < k; i++ {
		pop(&nums)
	}
	return pop(&nums)
}

func findKthLargest2(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

type IntHeap215 []int

func (h *IntHeap215) Len() int {
	return len(*h)
}
func (h *IntHeap215) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *IntHeap215) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *IntHeap215) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap215) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
