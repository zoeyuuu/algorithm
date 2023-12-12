package MonotoneStack

import "fmt"

// 84. 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/description/?envType=study-plan-v2&envId=top-100-liked
// 单调栈实现 (其实是考察每一个位置左右下一个更小的元素位置)

func Problem84() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea1(heights))
}

// 以每一个位置的高为基准 找到以这条高最大的宽 得到这条高对应的最大面积
// 寻找宽的时候 即找到左右比基准高更小的元素 可以在循环内扩充宽 但这样时间复杂度会到O(N^2)
// 用单调栈实现寻找每一个位置左右值更小的位置 下一个更小的位置 单调栈内部从栈顶到栈底应该满足严格递减
// 时间复杂度：O(N) 空间复杂度：O(N)
func largestRectangleArea1(heights []int) int {
	var monoStack []int
	lNext := make([]int, len(heights))
	rNext := make([]int, len(heights))
	// 下一个更小元素位置初始化为边界外
	for i := 0; i < len(heights); i++ {
		lNext[i] = -1
		rNext[i] = len(heights)
	}
	// 寻找右边下一个更小的元素
	for i := 0; i < len(heights); i++ {
		for len(monoStack) > 0 && heights[i] < heights[monoStack[len(monoStack)-1]] {
			rNext[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	// 寻找左边下一个更小的元素
	monoStack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[i] < heights[monoStack[len(monoStack)-1]] {
			lNext[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		maxArea = max(heights[i]*(rNext[i]-lNext[i]-1), maxArea)
	}
	return maxArea
}

// 思路：以不同的宽寻找最大的矩形面积 取最大值 （不太行
// 给定宽 最大面积 即寻找滑动窗口中最小值的最大值 用优先队列实现与p239相似
// 实现复杂 不能ac 87 / 98 个通过的测试用例 全1通过不了
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	ans := 0
	for i := 1; i <= n; i++ {
		area := maxAreaOfWidth(heights, i)
		fmt.Println("width：", i, "area:", area)
		if area > ans {
			ans = area
		}
	}
	return ans
}
func maxAreaOfWidth(nums []int, k int) int {
	n := len(nums)
	maxMin := 0
	var queue []int
	var push func(x int)
	push = func(x int) {
		// 注意队尾 > x 才弹出
		for len(queue) > 0 && queue[len(queue)-1] > x {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, x)
	}
	// 处理第一个窗口
	for i := 0; i < k; i++ {
		push(nums[i])
	}
	maxMin = queue[0]
	// 移动窗口
	for i := k; i < n; i++ {
		// Pop 若要删除的元素==队头 弹出队头
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		push(nums[i])
		maxMin = max(maxMin, queue[0])
	}
	return maxMin * k
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
