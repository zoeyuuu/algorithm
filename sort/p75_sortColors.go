package sort

import "fmt"

// 75. 颜色分类 medium 荷兰国旗问题
// https://leetcode.cn/problems/sort-colors/
// 双指针处理 第二个头尾指针更好理解  但...

func Problem75() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

// 双指针
// l1是排序好的0的后一个位置;l2是排序好的1的后一个位置
// nums[0:l1]==0;nums[l1:l2]==1
func sortColors(nums []int) {
	l1, l2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			// 与l2互换 l2++
			nums[l2], nums[i] = nums[i], nums[l2]
			l2++
		}
		if nums[i] == 0 {
			// 先将0安排到正确位置 分两种情况
			nums[l1], nums[i] = nums[i], nums[l1]
			if l1 < l2 { // 1.nums[l1]==1 此时nums[i]==1 需要进一步与l2交换
				nums[l2], nums[i] = nums[i], nums[l2]
			}
			// 2.l1==l2 也就是nums[l1]！=1 不需要进一步交换
			// l1,l2都需要++
			l1++
			l2++
		}
	}
}

// 双指针第二种写法 头尾指针p0之前全为0;p1之后全为2
// 大循环：i <= p1 遍历指针大于处理2的p1指针
// case nums[i] == 0:  swap(i,p0) i++
// case nums[i] == 1:  i++
// case nums[i] == 2:  swap(i,p1) 此时还要继续看交换完的nums[i]的值继续重复以上操作
// 总结：当nums[i] == 2需要进入循环交换i,p1直至！= 2 然后根据nums[i]=0/1 处理
func sortColors1(nums []int) {
	p0, p1 := 0, len(nums)-1
	for i := 0; i <= p1; i++ {
		for ; i <= p1 && nums[i] == 2; p1-- {
			nums[i], nums[p1] = nums[p1], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
