package heap

import (
	"container/heap"
	"sort"
)

// 直接内嵌结构体，这样可以省去 heap 的接口构造（Len()、Swap()）
type hp struct {
	sort.IntSlice
}

// Less 大顶堆
func (h *hp) Less(x, y int) bool {
	return h.IntSlice[x] > h.IntSlice[y]
}
func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *hp) Pop() interface{} {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return x
}
func lastStoneWeight(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.Pop(), q.Pop()
		if x > y {
			q.Push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
