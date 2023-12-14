package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/top-k-frequent-elements/
// 方法一： 先统计出现频率 直接使用sort.Slice排序 时间复杂度O(NlogN) 不满足时间复杂度要求 但能过 能过就行
// 方法二:  先统计出现频率 用小根堆排序 时间复杂度O(Nlogk)

func Problem347() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent1(nums, k))
}

// sort.Slice 排序
// 自己写的 能过
func topKFrequent1(nums []int, k int) []int {
	mp := map[int]int{}
	var list [][]int
	for _, v := range nums {
		mp[v]++
	}
	for key, value := range mp {
		list = append(list, []int{key, value})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] > list[j][1]
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = list[i][0]
	}
	return res
}

// 堆排序
// 小根堆 维持k个最高频出现的元素
// 时间复杂度： O(Nlogk)  每次堆操作需要O(logk)的时间
func topKFrequent2(nums []int, k int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range mp {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		// 最后弹出的是频率最高的 要反着来
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// IHeap 每个元素是[2]int{key, value}长度为2的数组
type IHeap [][2]int

func (h *IHeap) Len() int {
	return len(*h)
}

// Less 返回前K个高频 采用小顶堆 比较的值是(*h)[i][1]
func (h *IHeap) Less(i, j int) bool {
	return (*h)[i][1] < (*h)[j][1]
}
func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *IHeap) Push(x interface{}) {
	// 强制转换为长度为2的数组
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
