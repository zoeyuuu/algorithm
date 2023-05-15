package hashmap

// 两数之和
import "fmt"

func Problem1() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum2(nums, target))
}

// 哈希表法:遍历的过程中就查找target-x,并存储
func twoSum(nums []int, target int) []int {
	mp := map[int]int{} //键：数；值：下标
	for i, p := range nums {
		if j, ok := mp[target-p]; ok {
			return []int{i, j}
		}
		mp[p] = i
	}
	return nil
}

// 暴力
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
