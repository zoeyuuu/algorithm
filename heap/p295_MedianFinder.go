package heap

import (
	"container/heap"
	"fmt"
)

// 295. 数据流的中位数 hard 2023-04-24 26 hot100
// 用go实现蛮复杂的 考虑后期用java调库
// https://leetcode.cn/problems/find-median-from-data-stream/description/?envType=study-plan-v2&envId=top-100-liked
// 本题主要想实现在时间复杂度O(1)的情况下获得中位数 一般的方法会过不了 比如说sort
// 用两个堆(左边是大顶堆MaxHeap[较小部分的数]、右边小顶堆MinHeap[较大部分的数])维护 根据左右堆的大小进行添加操作和取中位数操作、
// 保证左边大顶堆MaxHeap的大小最多比MinHeap大小多1
// 时间复杂度：
// addNum:O(logn)，其中 n 为累计添加的数的数量。
// findMedian: O(1)
// 空间复杂度：O(n)，主要为优先队列的开销。
// m.maxHp.Push(num) ！！！第一次写错了 应该用heap.push这样每次push会自动调整堆 heap[0]才是最大元素

func Problem295() {
	median := Constructor()
	median.AddNum(1)
	median.AddNum(2)
	fmt.Println(median.FindMedian()) // 1.5
	median.AddNum(3)
	fmt.Println(median.FindMedian()) // 2

}

type MedianFinder struct {
	maxHp *MaxHeap
	minHp *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{&MaxHeap{}, &MinHeap{}}
}

// AddNum 根据MinHeap、MaxHeap大小来判断
func (m *MedianFinder) AddNum(num int) {
	// 此时应使maxHp.Len++
	if m.maxHp.Len() == m.minHp.Len() {
		// num与右边minHp的栈顶元素比较大小 确定num应该加入 minHp还是maxHp
		// 栈空或者num <= (*m.minHp)[0] ---> num ∈ maxHp 直接加入
		if m.minHp.Len() == 0 || num <= (*m.minHp)[0] {
			// m.maxHp.Push(num) ！！！第一次写错了 应该用heap.push这样每次push会自动调整堆 heap[0]才是最大元素
			heap.Push(m.maxHp, num)
		} else {
			// num>(*m.minHp)[0] ---> num ∈ minHp:
			// 将minHp的栈顶放入maxHp 然后将num加入minHp 保证maxHp.Len大于minHp.Len
			heap.Push(m.maxHp, heap.Pop(m.minHp))
			heap.Push(m.minHp, num)
		}
		// 此时应使minHp.Len++
	} else {
		// num ∈ minHp
		if num >= (*m.maxHp)[0] {
			heap.Push(m.minHp, num)
		} else {
			// num ∈ maxHp
			heap.Push(m.minHp, heap.Pop(m.maxHp))
			heap.Push(m.maxHp, num)
		}
	}
}

// FindMedian 根据大小顶对堆的大小返回
func (m *MedianFinder) FindMedian() float64 {
	// 大顶堆比小顶堆多1 -> 总长度是奇数 直接返回大顶堆堆顶元素
	if m.maxHp.Len() > m.minHp.Len() {
		return float64((*m.maxHp)[0])
	} else { // 左右堆大小相等 返回两个堆顶平均值
		return float64((*m.minHp)[0]+(*m.maxHp)[0]) / 2.0
	}
}

type MinHeap []int

func (m *MinHeap) Len() int {
	return len(*m)
}
func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}
func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() any {
	v := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return v
}

type MaxHeap []int

func (m *MaxHeap) Len() int {
	return len(*m)
}
func (m *MaxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}
func (m *MaxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *MaxHeap) Pop() any {
	v := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return v
}
