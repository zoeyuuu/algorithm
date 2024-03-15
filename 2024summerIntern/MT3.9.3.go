package main

import "fmt"

// 二维前缀和 消除两层循环
func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]int, n)
	// 读取数据
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		var numStr string
		// 一行数字没有空格 按字符串读取在挨个变成数字
		fmt.Scan(&numStr)
		for j := 0; j < n; j++ {
			matrix[i][j] = int(numStr[j] - '0')
		}
	}
	// 二维前缀和 长度n+1 下标存储 不用初始化
	pre := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			// 前缀和计算方式 两个矩形叠加减去重复再加上新元素（元素值注意是索引存储 要-1
			// 下标转换特别容易错！！
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	// 返回子矩阵的和 索引
	getSubMatirx := func(x1, y1, x2, y2 int) int {
		// [x2y2]-[x1-1y2]-[x2y1-1]+[x1y1] 索引分别+1
		return pre[x2+1][y2+1] - pre[x1][y2+1] - pre[x2+1][y1] + pre[x1][y1]
	}
	ans := make([]int, n)
	// 遍历所有位置
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// k+1 代表正方形长度
			for k := 1; k < n; k++ {
				// 越界检查
				if i+k >= n || j+k >= n {
					break
				}
				subMatrix := getSubMatirx(i, j, i+k, j+k)
				if subMatrix*2 == (k+1)*(k+1) {
					ans[k]++ // k+1存在索引k处 正好跟输出要求抵消
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}
