package main

import (
	"fmt"
	"math"
)

// 计算数组前后部分和的成绩的最大值 前后缀
// 过了25% bigint用了很久报错 int64可以？

func minProductOfSums(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0 // 如果数组长度小于2，无法进行划分
	}

	// 计算前缀和
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}

	totalSum := prefixSum[n-1]
	minProduct := math.MaxInt64

	// 遍历每个可能的划分点
	for i := 0; i < n-1; i++ {
		leftSum := prefixSum[i]
		rightSum := totalSum - leftSum
		product := leftSum * rightSum
		if product < minProduct {
			minProduct = product
		}
	}

	return minProduct
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Minimum product of sums:", minProductOfSums(arr))
}
