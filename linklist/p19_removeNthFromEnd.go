package linklist

func Problem19() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next.Next = &ListNode{Val: 6}
	printList(l1)
	printList(removeNthFromEnd(l1, 2))
}

// 方法一 O() O()
// 第一次遍历得到链表长度 第二次遍历到第L−n+1个节点 删除
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)

	//添加一个哑节点（dummy node），不需要对头节点进行单独判断
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

//方法二 用栈存储
