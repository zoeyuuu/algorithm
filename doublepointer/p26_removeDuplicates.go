package doublepointer

import "fmt"

func problem26() {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 3, 4, 5, 6}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

// 26. 删除有序数组中的重复项
// 双指针法

func removeDuplicates(nums []int) int {
	n := len(nums)
	i, j := 0, 1
	temp := nums[0]
	for ; i < n; i++ {
		for nums[i] != temp {
			temp = nums[i]
			nums[j] = temp
			j++
		}
	}
	return j
}
