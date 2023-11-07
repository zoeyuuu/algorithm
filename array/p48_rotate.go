package array

// 48 将图像 原地 顺时针旋转 90 度 meidum
// 正方形任意旋转可改为多重折叠
// 顺时针旋转90：先沿对角线反转矩阵，再沿竖中轴线反转矩阵；
// 顺时针旋转180：先沿横中轴线反转矩阵，再沿竖中轴线反转矩阵；
// 顺时针旋转270：先沿对角线反转矩阵，再沿横中轴线反转矩阵；

import "fmt"

func Problem48() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	matrix = rotate(matrix)
	printMatrix(matrix)
}

func rotate(matrix [][]int) [][]int {
	n := len(matrix)

	// 按对角线反转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 按竖中轴反转
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			//下标0和下标n-1交换
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	return matrix
}

/*
// 按横中轴线反转矩阵
	for i := 0; i < n/2 ; i++ {
	    for j := 0; j < n; j++ {
	        matrix[n-1-i][j], matrix[i][j] = matrix[i][j], matrix[n-1-i][j]
	    }
	}
*/

func printMatrix(matrix [][]int) {
	for _, nums := range matrix {
		for _, nums := range nums {
			fmt.Print(nums, " ")
		}
		fmt.Println()
	}
}
