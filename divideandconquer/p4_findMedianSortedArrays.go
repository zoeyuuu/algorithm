package divideandconquer

// 4. 寻找两个正序数组的中位数 hard
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。
import (
	"fmt"
)

func Problem4() {
	n1 := []int{1, 3, 5}
	n2 := []int{2, 4, 6, 7}
	fmt.Println(findMedianSortedArrays1(n1, n2)) //直接 O(M+N)时间复杂度不达要求 ac还行
	fmt.Println(findMedianSortedArrays2(n1, n2)) //分治法  O(log(M+N))
}

// 方法一 时间复杂度不达要求
// 合并两个数组后直接取中位数即可
// 时间复杂度 O(M+N) // not ok
func findMedianSortedArrays1(n1, n2 []int) float64 {
	res := merge(n1, n2)
	n := len(res)
	if n == 0 {
		return -1
	}
	if n%2 == 0 {
		return float64(res[n/2-1]+res[n/2]) / 2 //注意:1、数组下标要减一 2、先float64再/2
	}
	return float64(res[n/2]) //注意数组下标要减一
}

// 合并两个有序数组
func merge(n1, n2 []int) []int {
	l1, l2 := len(n1), len(n2)
	res := make([]int, 0, l1+l2)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if n1[i] <= n2[j] {
			res = append(res, n1[i])
			i++
		} else {
			res = append(res, n2[j])
			j++
		}
	}
	// 处理未合并部分
	if i < l1 {
		res = append(res, n1[i:]...)
	}
	if j < l2 {
		res = append(res, n2[j:]...)
	}
	return res
}

// 方法二
// 类二分查找分而治之的思路
// O(log(M+N)) // ok
// 这个函数的实现采用了分而治之的思路，每次递归将两个数组中的一部分舍弃：
// 基本思想是在两个数组中找第k/2大的元素 比较大小 谁更小则舍弃该数组前半部分所有元素 转而再剩下的两个数组中找k-k（舍弃）大的元素
// 从而将查找第 k 大的元素的问题转化为查找第 k-k1 或 k-k2 大的元素的问题，递归直到找到第 k 大的元素为止。

func findMedianSortedArrays2(n1, n2 []int) float64 {
	l1, l2 := len(n1), len(n2)
	if l1+l2 == 0 {
		return -1
	}
	if (l1+l2)%2 == 0 {
		//若长度为偶数 找到最中间的两个数
		v1 := findKth(n1, n2, (l1+l2)/2)
		v2 := findKth(n1, n2, (l1+l2)/2+1)
		return float64(v1+v2) / 2
	}
	return float64(findKth(n1, n2, (l1+l2)/2+1))
}

// 在两个有序数组中查找第 k 大的数
// 第 k 大的数即 nums 中索引为 k-1 的数，在舍弃区域比较值时都要 -1 处理
func findKth(n1, n2 []int, k int) int {
	l1, l2 := len(n1), len(n2)
	//避免数组长度的分类讨论 保证n1更短
	if l1 > l2 {
		l1, l2 = l2, l1
		n1, n2 = n2, n1
	}
	if l1 == 0 { //特殊情况
		return n2[k-1]
	}
	if k == 1 { //特殊情况
		return Min(n1[0], n2[0])
	}

	k1 := Min(k/2, l1) //考虑到l1可能大于k/2
	k2 := k - k1

	// 分治核心思想
	// 注意下标都要减一
	switch {
	case n1[k1-1] < n2[k2-1]:
		return findKth(n1[k1:], n2, k2)
	case n1[k1-1] > n2[k2-1]:
		return findKth(n1, n2[k2:], k1)
	default:
		return n1[k1-1]
	}
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
