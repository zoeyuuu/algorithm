package backtracking

import "fmt"

// 216. 组合总和 III medium
// https://leetcode.cn/problems/combination-sum-iii/description/
// 两种形式
// 1.p216  for遍历每一层是直接利用自然数 1~n 终止条件以len(path) == k ,k是targetLength
//  结果添加的时候path初始化不能有长度 path = append(path, i) 需要手动回溯path = path[:len(path)-1]
// 2.p17  for遍历每一层是map/数组的下标从0开始 终止条件以下标 == k
// 那么可以直接使用path[i] = value的形式添加 不需要手动回溯 每次循环直接覆盖上一次的值

func Problem216() {
	k := 3
	// n是target 可用数字是1~9
	n := 9
	fmt.Println(combinationSum3(k, n))
}
func combinationSum3(k, target int) [][]int {
	var res [][]int
	path := []int{}
	sum := 0
	var dfs func(startIndex, sum int)
	dfs = func(startIndex, sum int) {
		if len(path) == k {
			if sum == target {
				temp := make([]int, k)
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		// 剪枝1
		for i := startIndex; i <= 9-k+len(path)+1; i++ {
			// 剪枝2
			if sum+i > target {
				continue
			}
			path = append(path, i)
			dfs(i+1, sum+i)
			// 回溯
			path = path[:len(path)-1]
			// 错误：需要将sum+i作为递归参数 自动回溯 否则sum+=i dfs(i+1, sum)需要手动回溯：
			// sum -= i
		}
	}
	dfs(1, sum)
	return res
}
