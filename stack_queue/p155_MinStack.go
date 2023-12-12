package stack_queue

import "fmt"

// 155. 最小栈 meidum 2023-11-08 85
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈
// https://leetcode.cn/problems/min-stack/description/?envType=study-plan-v2&envId=top-100-liked
// 思路一：使用一个辅助栈 存储的是当前栈的最小值 与原始栈共进出
// 思路二：直接将栈结点换成node node中不仅存值 还存当前栈的最小值(见下)

func Problem155() {
	minStack := Constructor155()
	minStack.Push(3)
	minStack.Push(2)
	minStack.Push(1)
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	stack []node
}
type node struct {
	value, min int
}

func Constructor155() MinStack {
	return MinStack{}
}
func (st *MinStack) Push(val int) {
	if len(st.stack) == 0 {
		st.stack = append(st.stack, node{val, val})
	} else {
		min := st.stack[len(st.stack)-1].min
		if val < min {
			min = val
		}
		st.stack = append(st.stack, node{val, min})
	}
}
func (st *MinStack) Pop() {
	st.stack = st.stack[:len(st.stack)-1]
}
func (st *MinStack) Top() int {
	return st.stack[len(st.stack)-1].value
}
func (st *MinStack) GetMin() int {
	return st.stack[len(st.stack)-1].min
}
