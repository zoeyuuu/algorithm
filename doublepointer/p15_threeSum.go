package doublepointer

import (
	"fmt"
	"sort"
)

// 15 三数之和 medium
// https://leetcode.cn/problems/3sum/description/
// 返回元素组
// 暴力法和哈希（少一层循环）时间复杂的都高
//

func Problem15() {
	nums := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(nums))
}

// Problem3 :排序 + 双指针
/*
算法流程：
1、特判，对于数组长度 n，如果数组为 null或者数组长度小于 3，返回 []。
2、对数组进行排序。
3、遍历排序后数组：
    ·外层循环枚举a  （去重a）若nums[first]>0：因为已经排序好，所以后面不可能有三个数加和等于0，直接返回结果。
    · ***!!!内层循环枚举bc  (去重b) 使用采取双指针 将时间复杂度缩减到O(N)
         ·令左指针second=first+1，右指针 third=n−1，当 second<third时，执行循环：
         · second<third && nums[second]+nums[third] > target时 third--
         · nums[second]+nums[third] < target 时循环继续 second++
         · nums[second]+nums[third] = target 时循环继续 还是second++
           （此时third不变 变相利用second++ 实现了c的去重：nums[third]可能等于nums[third-1]但是third不变）
         · 如果指针重合 退出循环
时间复杂度：O(n^2)： 数组排序 O(NlogN)+（遍历数组 O(n)*双指针遍历 O(n)） =  O(n^2)
空间复杂度：O(1)
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	if nums == nil || n < 3 { // 特判
		return [][]int{}
	}
	sort.Ints(nums) //数组排序
	// 枚举 a (1.要去重 2. a如果大于target_sum 可以结束循环 因为数组递增）
	for first := 0; first < n && nums[first] <= 0; first++ {
		// 去重a
		if first > 0 && nums[first] == nums[first-1] { //first-1注意下标若为-1则报错
			continue
		}
		// 枚举 b、c
		// 枚举 bc的时候采取双指针 将时间复杂度缩减到O(N)
		target := 0 - nums[first]
		for second, third := first+1, n-1; second < third; second++ {
			// 去重b
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 注意下面这里是for 而不是if第一次写写错了
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合 退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
