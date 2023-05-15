package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印单链表
func printList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val, "->")
		p = p.Next
	}
	fmt.Println()
}

// 打印单链表
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}
