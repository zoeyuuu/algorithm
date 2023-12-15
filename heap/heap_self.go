package heap

// 用数组表示二叉树，对于下标i的结点:
// 父结点是(i-1)/2 左孩子i*2+1 右孩子i*2+2

// 自己写堆的几个函数 大小顶堆在down和up函数中区别
// 1.BuildMaxHeap(nums []int) 对于所有非叶子节点执行down操作
// 2.down(nums []int, i, n int) “向下冒泡” 调整堆
// 3.pop(nums *[]int) int 交换堆顶堆底元素 堆顶down操作 弹出堆底
// 4.push

// BuildMaxHeap 构建大顶堆
// 对于非叶子节点(index:n/2-1~0)执行“向下冒泡” down操作
func BuildMaxHeap(nums []int) {
	n := len(nums)
	// 最后一个非叶子节点(最后一个节点的父亲):((n-1)-1)/2 = n/2 - 1
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
}

// “向下冒泡”调整堆
// i:当前处理结点下标 n:堆长度
func down(nums []int, i, n int) {
	max := i                    // max标记自己和孩子中最大值结点下标
	left, right := 2*i+1, 2*i+2 // 左,右孩子
	// left < n 要保证不越界
	if left < n && nums[left] > nums[max] {
		max = left
	}
	if right < n && nums[right] > nums[max] {
		max = right
	}
	// 如果最大值是孩子 交换 并且递归处理下一个位置
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		down(nums, max, n)
	}
}

// pop 输出堆顶元素
// 需要更改切片长度 参数使用指针形式 函数返回之后仍然有效
func pop(nums *[]int) int {
	last := len(*nums) - 1
	// 交换堆顶堆底元素
	(*nums)[0], (*nums)[last] = (*nums)[last], (*nums)[0]
	// 新堆顶元素向下冒泡 注意此时堆长度-1=last
	down(*nums, 0, last)
	v := (*nums)[last]
	*nums = (*nums)[:last]
	return v
}

// push 向堆中插入元素
// 先将元素放在堆底 然后对其进行向上冒泡操作
func push(nums *[]int, value int) {
	*nums = append(*nums, value)
	up(*nums, len(*nums)-1)
}

// 向上冒泡调整堆
func up(nums []int, i int) {
	// 直到调整到堆顶
	for i > 0 {
		parent := (i - 1) / 2
		if nums[i] >= nums[parent] {
			break
		}
		// 交换当前节点与父节点
		nums[i], nums[parent] = nums[parent], nums[i]
		i = parent
	}
}
