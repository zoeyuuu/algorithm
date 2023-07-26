package greedy

import (
	"fmt"
)

// 134 canCompleteCircuit medium
// https://leetcode.cn/problems/gas-station/
// 算法整个题解 写的很好：https://leetcode.cn/problems/gas-station/solutions/2074636/dang-lao-si-ji-xue-hui-liao-tan-xin-suan-8s2g/
// 题目理解：gas/cost每个数据都要用到 不是考虑舍弃一个的问题 因为是问能不能回到原点
// 将gas-cost理解成经过加油站到下一个加油站的代价 这样题目即转换为对diff数组累加 然后保证暂存和要一直大于0知道并能够循环
// 是否有解：判断diff数组累加和是否>=0
// 解的唯一性：(可以结合上述链接的图像理解) 暂存和的最低点minSum
// 然后在diff数组累加的过程中 寻找（暂存和的最低点minSum） 从这里开始必定能保证

func Problem134() {
	//gas := []int{1, 2, 3, 4, 5}
	gas := []int{11, 4, 7, 1, 0}
	//cost := []int{3, 4, 5, 1, 2}
	cost := []int{2, 5, 5, 9, 1}
	fmt.Println(canCompleteCircuit(gas, cost))
}

// diff := []int{-2,-2,-2,3,3}
func canCompleteCircuit(gas, cost []int) int {
	n := len(gas)
	diff := make([]int, n)
	index := -1 //初始值为-1 最终都要+1
	sum := 0    //暂存和
	minSum := 0 //暂存和最小值
	//找到diff数组最小值下标index
	for i := 0; i < n; i++ {
		diff[i] = gas[i] - cost[i] //计算diff数组
		sum += diff[i]             //计算暂存和
		if sum < minSum {
			minSum = sum
			index = i
		}
	}
	if sum < 0 {
		return -1
	}
	return (index + 1) % n
}
