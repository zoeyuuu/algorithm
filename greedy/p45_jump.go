package greedy

// 45 跳跃游戏 II medium
// https://leetcode.cn/problems/jump-game-ii/

import "fmt"

func Problem45() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

// 自己做的
// 思路:贪心 找到离最后位置最远的且能跳到最后位置的index 然后找到离index最远的且能跳到index的下标 以此类推
// 先用一次循环将数组变为当前位置能到达的最远位置
// 从后往前 index为下标 直到index移到数组开始位置停止循环
// 循环内找到最远的能够到达index的位置index'并更新index 跳跃次数step++
func jump(nums []int) int {
	for i := 0; i < len(nums); i++ {
		nums[i] += i
	}
	index := len(nums) - 1 // 下标 从后往前寻找
	step := 0              //跳跃次数
	for index > 0 {
		for j := 0; j < index; j++ {
			if nums[j] >= index {
				index = j
				step++
			}
		}
	}
	return step
}

// 代码随想录算法:
// 遍历一次 next记录遍历到当前位置所能覆盖的最大距离
// cur记录当前一步所能达到的最远距离
// 如果当前位置==next 则表示需要跳一次 step++
func jump2(nums []int) int {
	n := len(nums)
	// 特殊情况
	if n == 1 {
		return 0
	}
	//cur:当前覆盖的最远距离下标
	//next:下一步覆盖的最远距离下标
	cur, next := 0, 0
	step := 0
	for i := 0; i < n-1; i++ {
		next = Max(nums[i]+i, next)
		if i == cur {
			//下一步能到达的最远距离
			cur = next
			step++
		}
	}
	return step
}
