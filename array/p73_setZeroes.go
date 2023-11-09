package array

// 73. 矩阵置零 meidum
// https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 思路：遍历两遍数组 第一遍用row、col两个map记录要设置0的下标 第二遍根据行列号置零
// 时间O(MN) 空间O(M+N)

func Problem73() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
}
func setZeroes(matrix [][]int) {
	row := make(map[int]bool)
	col := make(map[int]bool)
	for i, v := range matrix {
		for j, r := range v {
			if r == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, v := range matrix {
		for j, _ := range v {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
