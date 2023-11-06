package backtracking

import (
	"fmt"
)

// 491. 递增子序列 medium
// 与p78 p90求子集类似
// 1.类似求的是子集 需要在递归每一个结点存储结果（题目要求>=2）
// 2.递增子序列 遍历到nums[i]>=path[-1] --- if len(path) > 0 && nums[i] < path[len(path)-1] { continue }
// 3.去重 p90可以先对数组排序 然后通过每一层的该元素和上一个是否相同进行去重
// 考虑1,2,4,2,2 在遍历第二层的时候分别经过2,4,2,2; 到第二三个2时 used_map中已有元素2,4 此时跳过第二三个2
// 错误思路：
//用一个last变量维护每一层刚刚使用过的 进行去重 不行 因为不是有序的 还是需要used_map 比如上面第二层遍历到第二个2的时候last = 4
// 本题求的是递增序列 不能排序 故对于每一层进行去重的时候必须用map记录该层已使用过的元素 如果遍历元素在该层没使用过则继续

func Problem491() {
	nums := []int{1, 2, 1, 1, 1}
	fmt.Println(findSubsequences(nums))
}
func findSubsequences(nums []int) [][]int {
	n := len(nums)
	var path []int
	var ans [][]int
	var dfs func(index int)
	dfs = func(index int) {
		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
		// 每一层使用一个used数组去重
		// 写错：mp初始化必须用make 使用var声明的map变量会被初始化为nil
		// 在nil的map上执行赋值操作会导致"panic: assignment to entry in nil map"。
		used := make(map[int]bool)
		for i := index; i < n; i++ {
			//  每一层去重
			if used[nums[i]] {
				continue
			}
			// 递增要求
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			// 每一层加入path的元素 在used数组中记录
			used[nums[i]] = true
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
