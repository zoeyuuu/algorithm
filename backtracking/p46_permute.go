package backtracking

import "fmt"

// 46. 全排列 medium
// https://leetcode.cn/problems/permutations/description/
// 全排列 与组合问题的区别：1.每层都是从0开始搜索而不是startIndex; 2.used数组记录path已使用的元素
// 时间复杂度: O(n!)
// 空间复杂度: O(n)

func Problem46() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

// 处理的下标==len(nums)作为递归终止条件
func permute(nums []int) [][]int {
	// path初始化长度为0 递归条件与len(path)有关
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	// 数组used记录已使用的元素
	used := make([]bool, len(nums))
	// dfs的参数可以理解为是处理path的下标
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(cur + 1)
			// 回溯
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0)
	return res
}

// 形式2：dfs参数省略 使用path长度作为递归条件
func permute1(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	// 数组used记录已使用的元素
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			// 回溯
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
