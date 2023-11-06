package backtracking

import (
	"fmt"
	"sort"
)

// 90. 子集 II medium
// https://leetcode.cn/problems/subsets-ii/description/
// 与p78子集I相比 多了重复元素及结果去重
// p40(*)组合总数II去重思路一样 数组中有重复元素 要保证得到的子集不重复
// 1.先对数组进行排序 2.在每一层广度遍历的时候进行去重	if i > index && nums[i] == nums[i-1] { continue }

func Problem90() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var path []int
	var ans [][]int
	var dfs func(index int)
	dfs = func(index int) {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		for i := index; i < n; i++ {
			// 去重
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
