package greedy

// 另一种思路是用动态规划 再写一下
// 题目链接 https://leetcode.cn/problems/wiggle-subsequence/description/

// 题目要求是子序列 可以不是连续的 即可以删除一些元素得到子序列
// 但不需要删除一些不满足题目的要求的数 只需要遍历一遍统计即可
// 从头开始遍历 可以将算法抽象成 如果当前是上升趋势且之前是下降趋势 那么就可以视作一个摆动  子序列长度++ 否则继续遍历
// 最终的ans需要加一 因为最终的边界条件需要加一

import "fmt"

func Problem376() {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	fmt.Println(wiggleMaxLength(nums))
}

func wiggleMaxLength(nums []int) int {
	ans := 0     //总长度
	reverse := 0 //记录摆动状态 0代表初始 1代表下降 2代表上升
	for i := 1; i < len(nums); i++ {
		//之前不是下降趋势 现在是 ans++
		if reverse != 1 && nums[i] < nums[i-1] {
			ans++
			reverse = 1
		}
		if reverse != 2 && nums[i] > nums[i-1] {
			ans++
			reverse = 2
		}
	}
	return ans + 1
}
