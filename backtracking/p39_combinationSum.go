package backtracking

import "fmt"

// 39. 组合总和 meidum
// https://leetcode.cn/problems/combination-sum/description/
// candidates 中的 同一个 数字可以 无限制重复被选取

func Problem39() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	n := len(candidates)
	var path []int
	var dfs func(index, sum int)
	// 求的是组合 需要startIndex避免重复
	dfs = func(index, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			// 写错temp声明要有长度>=len(path)才能用copy
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := index; i < n; i++ {
			path = append(path, candidates[i])
			// 可重复利用：i不用+1
			dfs(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
