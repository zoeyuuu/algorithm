package greedy

import (
	"fmt"
	"sort"
)

func Problem455() {
	nums1 := []int{1, 2, 3} //胃口g
	nums2 := []int{1, 1}    //饼干s
	fmt.Println(findContentChildren(nums1, nums2))
}

// greedy 贪心算法
// 思路：将最大的饼干给胃口最大的
// 优先考虑胃口 以胃口为外层循环 不满足条件则饼干下标递减
// 需要先将数组进行排序 使用sort.Ints(g) 时间复杂度为O(nlogn)
func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	// 优先考虑胃口 不符合要求的时候 胃口-- 胃口在外层
	var result = 0
	index := len(s) - 1 //饼干下标 在if内循环--
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && g[i] <= s[index] {
			result++
			index--
		}
	}
	return result
}
