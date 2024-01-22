package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

// 1046. 最后一块石头的重量 easy
// https://leetcode.cn/problems/last-stone-weight/description/
// 方法一：最大堆 ；方法二 排序（模拟最大堆）

func Problem1046() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight1(stones))
	fmt.Println(lastStoneWeight2(stones))
}

// 方法一：最大堆实现 最大堆的定义
// 实现heap.Interface的接口
// 1.创建一个包含整数切片的结构体 2.实现heap.Interface接口的方法：
// Ctrl+I 输入interface(heap) 快速实现heap.Interface接口的5个方法 其中Push/Pop方法接收者类型为指针
// 3.创建一个空堆、初始化堆 4.堆中push和pop使用heap.Push(q, v)、heap.Pop(q)即可调用自己写的Push\Pop方法

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// Less 大顶堆 小顶堆h[i] < h[j]
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push / Pop 需修改接收者值 使用指针类型
func (h *IntHeap) Push(x any) {
	// x（interface{}类型）转换为int类型
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func lastStoneWeight1(stones []int) int {
	q := &IntHeap{} // 创建一个空的 IntHeap堆 ; q指针
	heap.Init(q)    // 初始化堆
	for _, v := range stones {
		heap.Push(q, v)
		fmt.Println(q)
	}
	for q.Len() > 1 {
		// 强制转换为int类型
		x, y := heap.Pop(q).(int), heap.Pop(q).(int)
		if x > y {
			heap.Push(q, (x - y))
		}
	}
	if q.Len() > 0 {
		return (*q)[0]
	}
	return 0
}

// 方法二：排序
func lastStoneWeight2(stones []int) int {
	n := len(stones)
	for n > 1 {
		// 每次循环都要排序
		sort.Ints(stones) // 递增排序
		x, y := stones[n-1], stones[n-2]
		if x == y {
			stones = stones[:n-2]
		} else {
			stones = append(stones[:n-2], x-y)
		}
		n = len(stones) // 更新切片长度
	}
	if n > 0 {
		return stones[0]
	}
	return 0
}

//    使用自定义排序函数对切片进行降序排序
//    sort.Slice(stones, func(i, j int) bool {
//        return stones[i] > stones[j]
//    })
