package doublepointer

import (
	"fmt"
	"sort"
)

// 15 三数之和
// threeSum1/threeSum2三重循环暴力解法能通过95%
// threeSum3使用双指针法

func Problem15() {
	nums := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum3(nums))
}

// 暴力解法 三个循环遍历 对三元组排序查找去重
// 结果： 308 / 312 个通过的测试用例 超出时间限制
// 改进 1、双指针法 2、去重查找使用哈希表(试了 效果不大) 3、三元组排序可以使用GO自带的sort包？
func threeSum1(nums []int) [][]int {
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					three := sortThree([]int{nums[i], nums[j], nums[k]})
					flag := true
					for _, p := range ans {
						if three[0] == p[0] && three[1] == p[1] && three[2] == p[2] {
							flag = false
						}
					}
					if flag {
						ans = append(ans, three)
					}
				}
			}
		}
	}
	return ans
}

// 暴力算法 查找去重的时候使用hashmap 提升不大
func threeSum2(nums []int) [][]int {
	var ans [][]int
	mp := map[[3]int]struct{}{} //切片不能直接比较 map的键可以存为数组
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					p := sortThree([]int{nums[i], nums[j], nums[k]})
					three := [3]int{p[0], p[1], p[2]} //转换为[3]int数组
					if _, ok := mp[three]; !ok {
						mp[three] = struct{}{}
						ans = append(ans, p)
					}
				}
			}
		}
	}
	return ans
}

func sortThree(nums []int) []int {
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	if nums[1] > nums[2] {
		nums[1], nums[2] = nums[2], nums[1]
	}
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	return nums
}

// Problem3 :排序 + 双指针
/*
算法流程：
1、特判，对于数组长度 n，如果数组为 null或者数组长度小于 3，返回 []。
2、对数组进行排序。
3、遍历排序后数组：
    ·若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于0，直接返回结果。
    ·对于重复元素：跳过，避免出现重复解
    ·令左指针L=i+1，右指针 R=n−1，当 L<R时，执行循环：
         ·当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。
          并同时将 L,R移到下一位置，寻找新的解
         ·若和大于 0，说明 nums[R]太大，R左移
         ·若和小于 0，说明 nums[L]太小，L右移
复杂度分析：
时间复杂度：O(n^2)，数组排序 O(NlogN)，遍历数组 O(n)，双指针遍历 O(n)，总体 O(n^2)
空间复杂度：O(1)
*/
func threeSum3(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	if nums == nil || n < 3 { // 特判
		return [][]int{}
	}
	sort.Ints(nums) //数组排序
	for first := 0; first < n && nums[first] <= 0; first++ {
		//保证first与上一次列举的不同以去重
		if first > 0 && nums[first] == nums[first-1] { //first-1注意下标若为-1则报错
			continue
		}
		for second, third := first+1, n-1; second < third; second++ {
			//保证second与上一次列举的不同以去重(注意不能与first相同)
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//注意下面这里是for 而不是if第一次写写错了
			for nums[first]+nums[second]+nums[third] > 0 && second < third {
				third--
			}
			if second == third {
				break
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
