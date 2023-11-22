package array

import "fmt"

// 118. 杨辉三角 easy 2023-09-15 9
// https://leetcode.cn/problems/pascals-triangle/description/?envType=study-plan-v2&envId=top-100-liked
// 逐行生成 每一行首尾等于1 其他值等于上一行错位两个数相加

func Problem118() {
	n := 3
	fmt.Println(generate(n))
}
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
