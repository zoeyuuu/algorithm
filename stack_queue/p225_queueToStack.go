package stack_queue

import "fmt"

// 225. 用队列实现栈 2022-10-30 29
// https://leetcode.cn/problems/implement-queue-using-stacks/description/

func Problem225() {
	stack := Constructor1()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Empty())
}

// 自己写的 用一个队列实现
// 需要pop的时候 将n-1个元素重新入队然后返回队首的元素 实现弹出的是最新压入的元素 (栈

type MyStack struct {
	queue []int //只允许在队首操作
}

func Constructor1() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.queue = append(s.queue, x)
}

func (s *MyStack) Pop() int {
	for i := 1; i <= len(s.queue)-1; i++ {
		x := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, x)
	}
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *MyStack) Top() int {
	v := s.Pop()
	s.Push(v)
	return v
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}
