package hashmap

import "fmt"

// 41. 缺失的第一个正数 hard
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。 (原地哈希 方法二)

func Problem41() {
	nums := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(nums))
}

// 方法一：一般哈希 如果没有空间复杂度限制 用这个
// 思路：给定长为N的数组，最小的正整数的范围只能在1~N+1之间(最极限的情况时1、2、3...N)
// 创建一个长为N+1的数组flag(0不考虑) 记录1~N的正整数是否出现 出现设置为true
// 遍历flag 返第一个没出现(false)的数
// 时间复杂度O(N) 空间复杂度O(N)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	flag := make([]bool, n+1)
	for i := 0; i < n; i++ {
		if nums[i] > 0 && nums[i] <= n {
			flag[nums[i]] = true
		}
	}
	for i := 1; i <= n; i++ {
		if flag[i] == false {
			return i
		}
	}
	return n + 1
}

// 使用数组本身作为哈希表（nums[i]=i+1代表i+1出现过） 不需要额外的空间复杂度
// 数组长度n，下标0~n-1用来表示1~n是否存在 注意下标比值小一
// 定义nums[i]=i+1代表i+1存在 遇到1<=nums[i]<=N  将(-1转换为下标)nums[i]-1下标的值改为nums[i]
// 同时还要继续处理原来nums[i]-1下标的值 避免覆盖 这里使用交换操作 将nums[i]-1下标的值交换到nums[i]然后while循环继续处理此处的值
// 时间复杂度O(N) 空间复杂度O(1)
func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 遇到nums[i]先检查是否在1~n内 若在检查对应下标的哈希值是否等于i+1
		// 如果不等于 那么交换nums[i]和nums[index]的值(防止覆盖) 继续处理新的nums[i]的值
		// nums[nums[i]-1] ？= nums[i]-1 + 1 保证哈希表值正确
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
