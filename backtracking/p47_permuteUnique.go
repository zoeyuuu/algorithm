package backtracking

import (
	"fmt"
	"sort"
)

func Problem47() {
	nums := []int{1, 1, 3}
	fmt.Println(permuteUnique(nums))
}
func permuteUnique(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	// 先排序
	sort.Ints(nums)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := 0; i < len(nums); i++ {
			// 去重 与p40层次遍历去重方法一致 添加used[i-1]==flase判断条件
			// ***在数层去重时used[i-1]==true代表上一层使用过 这种情况层级不用去重***
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				dfs(cur + 1)
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	dfs(0)
	return res
}
