package linklist

// 2. 两数相加 medium 2023-12-19 107
// 单链表
// https://leetcode.cn/problems/add-two-numbers/description/

func Problem1() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 9}

	result := addTwoNumbers(l1, l2)
	printList(result)
}

// 同时遍历两个链表 计算和及进位 存储到一个新链表(单独处理头结点
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	// 大循环 只要一个链表有值就继续
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 计算本位和进位
		value := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		// 链表为空单独处理
		if head == nil {
			head = &ListNode{Val: value}
			tail = head
		} else {
			tail.Next = &ListNode{Val: value}
			tail = tail.Next
		}
	}
	// 最后处理进位
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}
	return head
}
