package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}
func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *minHeap) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
		fmt.Scan(&arr[i][0])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i][1])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	hp := &minHeap{}
	heap.Init(hp)
	var maxSum, ans int64
	for i := 0; i < n; i++ {
		heap.Push(hp, arr[i][0])
		// 遍历i>k个之后 maxSum要减掉堆顶小的元素
		if i > k-1 {
			maxSum += int64(arr[i][0])
			maxSum -= int64(heap.Pop(hp).(int))
		} else {
			maxSum += int64(arr[i][0])
		}
		// 遍历i>=k个之后 计算答案
		if i >= k-1 {
			fmt.Println(arr[i][1], maxSum)
			ans = max(ans, int64(arr[i][1])*maxSum)
		}
	}
	fmt.Println(ans)
}
func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

/*
思路：对于每一个评论数研究 最大优秀度等于>=该评论数的点赞数之和*评论数
所以按照评论数排序 从大到小 用优先队列存储>=评论数的点赞数集合 然后取K个最大的点赞数
*/

/*小红在小红书上面发布了n篇笔记，其中第i筒笔记的点赞数量为ai,评论数为bi,。
现在小红准备选择k筒笔记作为“精选笔记合集”，合集的优秀程度为:所有笔记点赞数之和乘以评论数的最小值。
现在小红想知道，最终台集最大的优秀度是多少?
输入描述:
第一行输入两个正整数n,k,代表笔记的数量,以及小红准备选择的合集大小。第二行输入n个正整教ai,代表每简笔记的点赞数。
第三行输入n个正整数bi,代表每篇笔记的评论数。
输出描述:
—个正整数,代表最终最大的优秀度。
输入样例：
4 2
1 2 3 4
3 4 2 1
输出样例：
10
*/
