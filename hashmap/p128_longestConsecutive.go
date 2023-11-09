package hashmap

import (
	"fmt"
	"sort"
)

// 128. 最长连续序列 medium
// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked
// 方法一：先排序时O(nlogn) 空O(1) 方法二：哈希表 时O(n) 空O(n)
// 方法三：并查集 暂时忽略

func Problem128() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive1(nums))
	fmt.Println(longestConsecutive2(nums))
}

// 方法二：哈希表
// 先建立所有元素的哈希表 查找时间复杂度降到O(1)
// 对于map中的每一个元素 查找之后的num+1,num+2...是否存在
// 连续长度cur -num +1
// **每次查找num时 先查看num-1是否在map中 否则跳过 避免了重复查找 每个元素只被遍历一次 时间复杂度降到O(N)
func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	sum := 1
	for num := range mp {
		curLen := 1
		// 检查num-1是否在map中 否则跳过
		if !mp[num-1] {
			for mp[num+1] {
				num++
				curLen++
			}
		}
		// 统计长度
		if curLen > sum {
			sum = curLen
		}
	}
	return sum
}

// 方法一：先排序 然后每次遍历到num[i],看下一个num[i]+1在不在
// 时间空间均超过95% 但排序理论时间复杂度O(nlogn) 没有满足题目O(n)
func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	sum := 1
	maxSum := 1
	for i := 1; i < len(nums); i++ {
		// 相等继续
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] == nums[i-1]+1 { //numn[i]+1存在 sum++ 更新maxSum
			sum++
			maxSum = max(sum, maxSum)
		} else { //numn[i]+1不存在 sum重置
			sum = 1
		}
	}
	return maxSum
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
