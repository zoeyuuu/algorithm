package hashmap

import "fmt"

// 982 按位与为零的三元组
// 暴力算法时间复杂度O(n^3)
// 优化：预处理所有 nums[i]&nums[j]的出现次数 用哈希表/数组存储
// 优化后时间复杂度O(n^2+C*n) 其中C为存储每一种nums[i]&nums[j]出现的次数数组/哈希表的长度 n为数组nums的长度
// 空间复杂度 O(C)

func Problem982() {
	nums := []int{2, 1, 3}
	fmt.Println(countTriplets(nums))
}
func countTriplets(nums []int) int {
	cnt := make(map[int]int) //哈希表存储每一种nums[i]&nums[j]及其出现的次数
	for _, x := range nums {
		for _, y := range nums {
			cnt[x&y]++
		}
	}
	sum := 0
	for x, c := range cnt {
		for _, y := range nums {
			if x&y == 0 {
				sum += c
			}
		}
	}
	return sum
}
