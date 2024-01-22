package linklist

func Problem24() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next = &ListNode{Val: 6}
	printList(swapPairs(l1))
}

// 涉及到两个元素之间的交换 其实涉及到四个结点 三次操作
// pre,cur,next
func swapPairs(head *ListNode) *ListNode {
	// 建立虚结点 可以直接用head进行遍历
	dummy := &ListNode{0, head}
	pre := dummy
	// 循环条件时是cur!= nil && cur.Next !=nil 因为要交换两个结点
	for head != nil && head.Next != nil {
		// 存储第四个结点
		next := head.Next.Next
		// 三次操作
		pre.Next = head.Next
		pre.Next.Next = head
		head.Next = next
		// 移动双指针(注意是调整之后的顺序
		pre = head
		head = next
	}
	return dummy.Next
}
