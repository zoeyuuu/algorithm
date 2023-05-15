package linklist

// 2. 两数相加
// 单链表

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

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var p, head *ListNode        //这样声明初始条件head==nil
	carry := 0                   //进位
	for l1 != nil || l2 != nil { //大循环
		n1, n2 := 0, 0 //每个结点的值
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sum %= 10
		if head == nil { //判断头节点是否为空
			head = &ListNode{Val: sum}
			p = head
		} else {
			p.Next = &ListNode{Val: sum}
			p = p.Next
		}

	}
	if carry > 0 { //最后有进位需要添加新节点
		p.Next = &ListNode{Val: carry}
	}
	return head
}
