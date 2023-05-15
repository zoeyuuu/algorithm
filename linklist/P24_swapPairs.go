package linklist

func Problem24() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next = &ListNode{Val: 6}
	printList(swapPairs(l1))
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	// 删除过程想清楚 涉及四个节点
	for p := dummy; p.Next != nil && p.Next.Next != nil; {

	}
	return dummy.Next
}
