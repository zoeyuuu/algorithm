package greedy

import "fmt"

// 55 跳跃游戏 medium
// 其实跳几步无所谓，关键在于可跳的覆盖范围！
// 每次移动取最大跳跃步数（得到最大的覆盖范围），每移动一个单位，就更新最大覆盖范围。
// 贪心算法局部最优解：每次取最大跳跃步数（取最大覆盖范围），整体最优解：最后得到整体最大覆盖范围，看是否能到终点。

func Problem55() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	n := len(nums)
	maxIndex := 0
	temp := 0
	for i := 0; i <= maxIndex && i < n; i++ {
		// 第一次错写成了temp = maxIndex + nums[i]
		temp = i + nums[i]
		maxIndex = max(maxIndex, temp)
	}
	if maxIndex >= n-1 {
		return true
	}
	return false
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
