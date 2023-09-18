package array

import "fmt"

// 54. 螺旋矩阵I meidum (p59 II)
// https://leetcode.cn/problems/spiral-matrix/description/
// 这里的方法不需要记录已经走过的路径，所以执行用时和内存消耗都相对较小
// 首先设定上下左右边界
// 其次向右移动到最右，此时第一行因为已经使用过了，可以将其从图中删去，体现在代码中就是重新定义上边界
// 判断若重新定义后，上下边界交错，表明螺旋矩阵遍历结束，跳出循环，返回答案
// 若上下边界不交错(等于时继续)，则遍历还未结束，接着向下向左向上移动，操作过程与第一，二步同理
// 不断循环以上步骤，直到某两条边界交错，跳出循环，返回答案

func Problem54() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}

// 四个方向循环/边界改变 上下界和判断条件要注意
func spiralOrder(matrix [][]int) (res []int) {
	// 空矩阵返回
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	upper, right, down, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	for {
		// 上面一行
		for i := left; i <= right; i++ {
			res = append(res, matrix[upper][i])
		}
		// 重新设定上边界+越界判断
		if upper++; upper > down {
			break
		}
		// 右边一列
		for i := upper; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		if right--; right < left {
			break
		}
		//下面一行
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		if down--; down < upper {
			break
		}
		// 左边一列
		for i := down; i >= upper; i-- {
			res = append(res, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}
	return res
}
