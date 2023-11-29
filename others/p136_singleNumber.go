package others

import "fmt"

// 136. 只出现一次的数字 easy 2023-09-22 45
// 要求线性时间复杂度，常数级空间空间复杂度
// 使用位运算(异或⊕,golang中符号'^')
// 异或运算性质 1.a⊕0=a 2.a⊕a=0 3.满足交换律、结合律

func Problem136() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

// 所有元素异或运算 满足交换律 出现两次的元素异或得到0 一次元素与0异或得到自身
func singleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}
