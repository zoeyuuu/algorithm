package linklist

func reorderList(head *ListNode) {
	mid := middleNode(head)
	head1, head2 := head, reverseList(mid.Next)
	mid.Next = nil
	mergeLists(head1, head2)
}
func mergeLists(head1, head2 *ListNode) *ListNode {
	// dummy一定要:=初始化
	dummy := &ListNode{}
	cur := dummy
	for head1 != nil {
		cur.Next = head1
		head1 = head1.Next
		cur = cur.Next
		cur.Next = head2
		// 防止2比1短 得到空指针
		if head2 == nil {
			continue
		}
		head2 = head2.Next
		cur = cur.Next
	}
	return dummy.Next
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	// 判断条件 fast != nil && fast.Next != nil p876 找到偶数的是第二个重点
	// 本题偶数要找到前一个中点mid (mid.next得到第二个表头;mid.Next置空) 判断条件改：
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
