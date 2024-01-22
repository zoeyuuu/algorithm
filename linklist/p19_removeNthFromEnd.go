package linklist

// 19. 删除链表的倒数第 N 个结点 medium 2023-12-18 133
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/

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

// 一次遍历 快慢指针法
// 要找到倒数第n个结点 (到倒数第n+1个节点进行删除 pre指针
// pre指向dummy 判断一下+1-1 pre少走n步 即可到达 倒数第n+1个节点进行删除
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	i := 1
	for head != nil {
		head = head.Next
		// pre多走n步到达要删除的前一个节点的位置
		if i > n {
			pre = pre.Next
		}
		i++
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
