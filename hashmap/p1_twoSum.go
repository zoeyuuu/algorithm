package hashmap

import "fmt"

// 1 两数之和 easy
// https://leetcode.cn/problems/two-sum/description/

func Problem1() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
}

// 哈希表法:遍历的过程中就查找target-x,并存储
// 使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N)降低到 O(1)
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func twoSum1(nums []int, target int) []int {
	mp := map[int]int{} //键：数字 值：下标 查找map查找的是键
	for i, p := range nums {
		if j, ok := mp[target-p]; ok {
			return []int{i, j}
		}
		mp[p] = i
	}
	return nil
}

// 暴力
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func twoSum2(nums []int, target int) []int {
	for i, p := range nums {
		for j := i + 1; j < len(nums); j++ {
			if p+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
