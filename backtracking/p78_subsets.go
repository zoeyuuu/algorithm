package backtracking

import "fmt"

// 78. 子集 medium
// https://leetcode.cn/problems/subsets/description/
// 与p39组合不同 这里求的是子集 组合遇到递归所有的叶子节点停止 子集遇到每一个节点都返回
// 没有index == n才记录path的条件 每一个结点都记录; i=n 的时候自动返回
// 题目本身元素就是不重复的 不能包含重复的（顺序不重要） 所以需要递归函数需要index避免重复

func Problem78() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
func subsets(nums []int) [][]int {
	n := len(nums)
	var path []int
	var ans [][]int
	var dfs func(index int)
	dfs = func(index int) {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		for i := index; i < n; i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
