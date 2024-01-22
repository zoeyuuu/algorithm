package divideandconquer

import "fmt"

// 1901. 寻找峰值 II 每日一题
// https://leetcode.cn/problems/find-a-peak-element-ii/
// 二分法 题目假设峰值一定存在 相零格子值不相等
// 要求时间复杂度  O(mlog(n)) 或 O(nlog(m)) ----> 二分法
// 先找某一行的最大值 然后与上下格子比较 如果大于上下元素nums[i-1][j] < nums[i][j] > nums[i+1][j] 则找到
// 如果nums[i-1][j] < nums[i][j] < nums[i+1][j] 可以把范围缩小到i+1行之后 -> low = medium+1
// 同理nums[i-1][j] > nums[i][j] > nums[i+1][j] 可以把范围缩小到i-1行之前 -> high = medium-1
// 时间复杂度： O(mlog(n)) 每一行找最大值O(n) 列方向Olog(n)

func Problem1901() {
	mat := [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}
	fmt.Println(findPeakGrid(mat))
}
func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	low, high := 0, m-1
	for low <= high {
		i := low + (high-low)/2
		j := maxOfRow(mat[i])
		if i-1 >= 0 && mat[i][j] < mat[i-1][j] {
			high = i - 1
			// 出错！一定要continue 不执行之后语句
			continue
		}
		if i+1 < m && mat[i][j] < mat[i+1][j] {
			low = i + 1
			continue
		}
		// nums[i-1][j] < nums[i][j] > nums[i+1][j]
		return []int{i, j}
	}
	return nil //不可能
}
func maxOfRow(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}
