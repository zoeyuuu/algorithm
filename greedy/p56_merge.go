package greedy

import (
	"fmt"
	"sort"
)

// 56. 合并区间 medium
// 系列题目：p435 p452
// 由于是要合并区间 区间是从左端到右端 所以以左端点排序比较好
// https://leetcode.cn/problems/merge-intervals/

func Problem56() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for _, p := range intervals {
		// 左端点大于right 返回一个区间
		if p[0] > right {
			res = append(res, []int{left, right})
			left, right = p[0], p[1]
		} else {
			// 右端点更新为更大值
			right = Max(p[1], right)
		}
	}
	// 最后一个区间放入
	res = append(res, []int{left, right})
	return res
}
