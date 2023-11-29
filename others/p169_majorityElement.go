package others

import (
	"fmt"
	"math"
	"sort"
)

// 169. 多数元素 easy 2023-03-19 61
// https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-100-liked

func Problem169() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(nums))
}

// 方法一 哈希表统计
// 时间复杂度O(n)
func majorityElement1(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	for key, value := range mp {
		if value > len(nums)/2 {
			return key
		}
	}
	return math.MaxInt32
}

// 方法二 排序后在第2/n的元素必为众数
// 时间复杂度O(logn)
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票法
// 票数初始化为0 候选人初始化为nums[0] 当票数==0时 更换候选人
// 遇到相同的则 票数 + 1，遇到不同的则 票数 - 1
// 时间复杂度O(n)
func majorityElement3(nums []int) int {
	candidate := math.MaxInt32
	count := 0
	for i := range nums {
		if count == 0 {
			candidate = nums[i]
			count++
		} else {
			if nums[i] == candidate {
				count++
			} else {
				count--
			}
		}
	}
	return candidate
}
