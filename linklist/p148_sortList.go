package main

import "sort"

// 现存 再排序
func sortList(head *ListNode) *ListNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	dummy := &ListNode{Val: 0}
	cur := dummy
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		node := &ListNode{Val: arr[i]}
		cur.Next = node
		cur = node
	}
	return dummy.Next
}
