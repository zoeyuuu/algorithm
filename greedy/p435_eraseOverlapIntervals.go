package greedy

// 435 无重叠区间 medium
// 系列题目：p56 p435 p452
// https://leetcode.cn/problems/non-overlapping-intervals/description/
// 其实是找最大不重叠子空间 （给定一系列会议得到预定会议的最大数量）
// 总空间-最大不重叠子空间=需删除的重复空间
// 按照右边界递增排序 选择右边界最小的子空间作为第一个最大不重叠子空间 理由是右侧最小使得与后续空间重叠最不可能
// 删除与之重叠的子空间 再选一个接下来右边界最小的子空间 最后得到最大不重叠子空间
// 与p452思路类似 但是边界情况些许不同

import (
	"fmt"
	"sort"
)

func Problem435() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	right := intervals[0][1]
	nonOverlap := 1
	for _, p := range intervals {
		// 注意边界应为大于等于
		if p[0] >= right {
			nonOverlap++
			right = p[1]
		}
	}
	return len(intervals) - nonOverlap
}
