package greedy

import (
	"fmt"
	"sort"
)

// 452 findMinArrowShots medium
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/description/
// 贪心思想1：对于一个气球 如果可以从右边界射出 这样可以保证能够同时射爆很多气球 （数组按照右边界递增排序）
// 贪心思想2：按照右边界递增排序之后 从左到右遍历 如果当前的箭无法戳爆该气球 下一只箭即为该气球的右边界
// 初始化： 按照气球右边界递增排序后 第一个气球（不能被上一个气球戳爆）的右边界作为第一支箭的位置 箭的数量初始化为1
// 其实需要用一个数组记录该气球是否被戳爆 但是在循环过程中只要p[0]<pos 就代表该气球可以pos的箭戳爆 所以不用数组记录

func Problem452() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}

// 有点难想
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			ans++
			maxRight = p[1]
		}
	}
	return ans
}
