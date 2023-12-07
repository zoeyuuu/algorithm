package stack_queue

import "fmt"

// 232. 用栈实现队列 2023-10-26 120
// https://leetcode.cn/problems/implement-queue-using-stacks/description/
// 思路：双栈
// 将一个栈当作输入栈，用于压入 push 传入的数据；另一个栈当作输出栈，用于 pop 和 peek 操作。
// 每次 pop 或 peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序。

func Problem232() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}

type MyQueue struct {
	stIn, stOut []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.stIn = append(q.stIn, x)
}
func (q *MyQueue) Pop() int {
	// 若输出栈为空 将输入栈元素弹出并压入输出栈
	if len(q.stOut) == 0 {
		for len(q.stIn) > 0 {
			x := q.stIn[len(q.stIn)-1]
			q.stOut = append(q.stOut, x)
			q.stIn = q.stIn[:len(q.stIn)-1]
		}
	}
	v := q.stOut[len(q.stOut)-1]
	q.stOut = q.stOut[:len(q.stOut)-1]
	return v
}
func (q *MyQueue) Peek() int {
	v := q.Pop()
	q.stOut = append(q.stOut, v)
	return v
}
func (q *MyQueue) Empty() bool {
	return len(q.stIn) == 0 && len(q.stOut) == 0
}
