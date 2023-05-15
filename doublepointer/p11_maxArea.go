package doublepointer

import "fmt"

// 11. 盛最多水的容器 medium
// 给你 n 个非负整数 a1，a2，…，an，每个数代表坐标中的一个点 (i, ai) 。
// 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

func Problem11() {
	fmt.Println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// 方法一 暴力算法 测试用例55/61
func maxArea1(height []int) int {
	ans := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			ans = max(ans, (j-i)*min(height[i], height[j]))
		}
	}
	return ans
}

// 方法二 双指针法
// 思路： 头尾指针 向内移动值较短的指针
// 正确性：面积等于宽度*高度，向内移动高度较短的指针i 相当于排除了较短指针i所在位置所有情况
// 对于较短指针i所在位置所有情况：高度取决于最短的高度 h<=hi 宽度只会更短 所以面积必然小于该情况 可以全部移除
func maxArea2(height []int) int {
	n := len(height)
	ans := 0
	for left, right := 0, n-1; left < right; {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
