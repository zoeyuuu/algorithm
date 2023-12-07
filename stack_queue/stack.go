package stack_queue

import "fmt"

type Stack struct {
	stack []interface{}
}

// Push 元素x入栈
func (s *Stack) Push(x interface{}) {
	s.stack = append(s.stack, x)
}

// Pop 移除栈顶元素
func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}

// Top 返回栈顶元素
func (s *Stack) Top() interface{} {
	return s.stack[len(s.stack)-1]
}

// Empty 判断队列是否为空
func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}
func StackExample() {
	myStack := Stack{}
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
}
