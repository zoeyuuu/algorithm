package stack_queue

import "fmt"

// Queue 定义队列结构体
type Queue struct {
	queue []interface{}
}

// Push 将一个元素放入队列的尾部
func (q *Queue) Push(x interface{}) {
	q.queue = append(q.queue, x)
}

// Pop 移除队首元素
func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

// Peek 返回队首元素
func (q *Queue) Peek() interface{} {
	return q.queue[0]
}

// Empty 判断队列是否为空
func (q *Queue) Empty() bool {
	return len(q.queue) == 0
}
func QueueExample() {
	myQueue := Queue{}
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Empty())
}
