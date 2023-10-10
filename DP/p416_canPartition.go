package DP

import "fmt"

//******  0/1背包  ******
// 416 分割等和子集 medium
// https://leetcode.cn/problems/partition-equal-subset-sum/description/

// 利用0/1背包思想：原来dp[j]代表的是容量为j的背包，所背的最大物品价值
// （本题也是每个数字只能选一次 价值和重量等价）迁移到本题：dp[j] = j时表示容量为j的正好能装dp[j]的重量（数字和）
// 本题即看dp[target] =? sum/2
// 01背包（一维数组）的递推公式为：dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
// 本题 dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
// 初始化：dp[0]=0，value全>0则dp[i] = 0(否则dp[i]=负无穷)
// 遍历顺序：外层：物品遍历 内层：背包容量（倒序遍历！！！）

func Problem413() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
func canPartition(nums []int) bool {
	sum := 0
	for _, p := range nums {
		sum += p
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	// 物品遍历
	for i := 0; i < len(nums); i++ {
		// 倒序遍历背包容量
		// j: target~nums[i] j--
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return target == dp[target]
}
