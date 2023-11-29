package others

import "fmt"

// 69. x 的平方根 2023-11-29 110
// https://leetcode.cn/problems/sqrtx/description/
// 不允许使用任何内置指数函数和算符 计算平方根

func Problem69() {
	x := 99999
	fmt.Println(mySqrt(x))
}

// 二分查找 找到一个区间使得low^2 <= x且high^2 > x 返回low
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	low, high := 1, x
	// 保证low^2 <= x且high^2 > x
	// 最终状态 low+1 = high 答案是low
	for low+1 < high { // 根据终止条件确定for循环的条件
		mid := (low + high) / 2
		if mid*mid > x {
			high = mid //保证high^2 > x
		} else if mid*mid == x {
			return mid
		} else {
			low = mid
		}
	}
	return low
}
