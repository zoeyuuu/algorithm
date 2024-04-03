package stack_queue

import "fmt"

// 239. 滑动窗口最大值 hard
// 2023-09-26 98
// 单调队列 (用于求解滑动窗口最大/最小值问题)
// 不能使用最大堆的原因：最大堆可以返回最大元素 但无法移除滑动窗口移动的元素
// 设计单调队列的时候，pop，和push操作要保持如下规则： ***很重要
// pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
// push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
// 保持如上规则，每次窗口移动的时候，只要问que.front()就可以返回当前窗口的最大值。
// 时间复杂度：O(n)，其中 n 是数组 nums的长度。每一个下标恰好被放入队列一次，并且最多被弹出队列一次，因此时间复杂度为 O(n)。
// 空间复杂度：O(k) k是窗口大小
func Problem239() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	var queue []int
	// 总共n-k+1个答案
	res := make([]int, 0, n-k+1)
	var push func(x int)
	push = func(x int) {
		// 注意队尾 < x 才弹出
		for len(queue) > 0 && queue[len(queue)-1] < x {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, x)
	}
	// 处理第一个窗口
	for i := 0; i < k; i++ {
		push(nums[i])
	}
	res = append(res, queue[0])
	// 移动窗口
	for i := k; i < n; i++ {
		// Pop 若要删除的元素==队头 弹出队头
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		push(nums[i])
		res = append(res, queue[0])
	}
	return res
}
