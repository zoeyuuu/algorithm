package greedy

import (
	"fmt"
	"sort"
)

// 135 candy hard
// https://leetcode.cn/problems/candy/description/

func Problem135() {
	nums := []int{1, 2, 87, 87, 87, 2, 1}
	fmt.Println(candy1(nums))
}

// 随想录方法
// 如果在考虑局部的时候想两边兼顾，就会顾此失彼。
// 那么本题我采用了两次贪心的策略：
// 一次是从左到右遍历，只比较右边孩子评分比左边大的情况。
// 一次是从右到左遍历，只比较左边孩子评分比右边大的情况。
// 这样从局部最优推出了全局最优，即：相邻的孩子中，评分高的孩子获得更多的糖果。

func candy1(ratings []int) int {
	/**先确定一边，再确定另外一边
	    1.先从左到右，当右边的大于左边的就加1
	    2.再从右到左，当左边的大于右边的就再加1
	**/
	need := make([]int, len(ratings))
	sum := 0
	// 初始化(每个人至少一个糖果)
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	// 1.先从左到右，当右边的大于左边的 右边=左边+1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边且need不大于右边 左边=右边+1
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && need[i-1] <= need[i] {
			need[i-1] = need[i] + 1
		}
	}
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}

// 自己做法
// 之前尝试过一遍循环 不管是跟前面比较或是跟前后比较 后面的candies增改始终有可能推翻已经获得的大小条件 顾不可行
// 所以想到对数组由小到大进行candies的发放 先得到rating值由小到大的indices下标数组
// 然后由小到大对candies[indices[i]]的值进行累加(注意边界条件) 使之满足大于两边的条件
// 能通过 但是时间和空间效果都不佳
func candy2(ratings []int) int {
	n := len(ratings)
	// 初始化下标数组
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	// 得到ratings由小到大排序的下标数组
	sort.Slice(indices, func(i, j int) bool {
		return ratings[indices[i]] < ratings[indices[j]]
	})
	// 初始化candies
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	for i := range indices {
		// 排序后indices[i]下标的ratings[indices[i]]只会>=两边的值 只需考虑>的情况
		// 不是初始位置&& rating值大于前一个位置 需保证candies[indices[i]] > candies[indices[i]-1]
		if indices[i] != 0 && ratings[indices[i]] > ratings[indices[i]-1] {
			for candies[indices[i]] <= candies[indices[i]-1] { // <=
				candies[indices[i]]++
			}
		}
		// 不是最后位置&& rating值大于后一个位置 需保证candies[indices[i]] > candies[indices[i]+1]
		if indices[i] != n-1 && ratings[indices[i]] > ratings[indices[i]+1] {
			for candies[indices[i]] <= candies[indices[i]+1] {
				candies[indices[i]]++
			}
		}
	}
	sum := 0
	for i := range candies {
		sum += candies[i]
	}
	return sum
}
