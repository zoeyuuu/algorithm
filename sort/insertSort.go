package sort

func insertSort(nums []int, low, high int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		j := i - 1
		tmp := nums[i]                // 记录插入的值
		for j >= 0 && nums[j] > tmp { //左边比右边大
			nums[j+1] = nums[j] //右移1位
			j--                 //扫描前一个数
		}
		nums[j+1] = tmp // j+1为最后插入的值
	}
}
