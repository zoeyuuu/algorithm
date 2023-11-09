package hashmap

// 560. 和为 K 的子数组  meidum
// https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked
// 哈希表 前缀和

func Problem560() {
	nums := []int{1, 1, 1}
	k := 2
	subarraySum1(nums, k)
	subarraySum2(nums, k)
}

// 连续和可以用前缀和相减得到
// 前缀和+哈希表 时间复杂度O(n) 空间复杂度O(n)
// 遍历一遍 建立前缀和的map 键是前缀和 值是前缀和的个数
// 遍历的过程中就查找 前缀和为pre-k的个数
// k = pre-preI  preI = pre-k
func subarraySum1(nums []int, k int) int {
	mp := map[int]int{}
	// 前缀为0的个数为1
	mp[0] = 1
	pre, count := 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if value, ok := mp[pre-k]; ok {
			count += value
		}
		mp[pre] += 1
	}
	return count
}

// 暴力解法 时间复杂度O(n^2)
func subarraySum2(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
