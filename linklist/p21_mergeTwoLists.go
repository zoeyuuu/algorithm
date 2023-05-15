package linklist

// 21. 合并两个有序链表 easy

func Problem21() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 7}
	l2.Next.Next = &ListNode{Val: 9}

	printList(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy //用一个哑节点来处理nil问题
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}
