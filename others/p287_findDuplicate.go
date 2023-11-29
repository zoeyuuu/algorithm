package others

import (
	"fmt"
	"math"
)

// 287. 寻找重复数 2023-05-15 25
// https://leetcode.cn/problems/find-the-duplicate-number/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 要求不修改 数组 nums 且只用常量级 O(1) 的额外空间

func Problem287() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate1(nums))
}

// 二分查找 ：每一次猜一个数mid，然后遍历整个输入数组，进而缩小搜索区间，最后确定重复的是哪个数
// 根据抽屉原理：每次统计原始数组left~right中 小于等于 mid 的元素的个数 count
// 如果 count 严格大于 mid,重复元素就在区间 [left..mid] 里
// 否则，重复元素就在区间 [mid + 1..right] 里
// 直到left==right 找到最终的重复元素left
// 时间复杂度：O(NlogN)
// 空间复杂度O(1) 时间换空间
func findDuplicate1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			// 下一轮搜索的区间 [left..mid]
			right = mid
		} else {
			// 下一轮搜索的区间 [mid+1..right]
			left = mid + 1
		}
	}
	return left
}

// 基本解法 空间复杂度O(n)
func findDuplicate2(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	for key, value := range mp {
		if value >= 2 {
			return key
		}
	}
	return math.MaxInt32
}
