package linklist

func Problem203() {

}

// 203. 移除链表元素 easy 2021-08-20 4
// https://leetcode.cn/problems/remove-linked-list-elements/description/
// 使用一个dummy统一链表遍历
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	// p表示处理的前一个结点
	for p := dummy; p.Next != nil; {
		if p.Next.Val == val {
			// 不用移动p 因为p代表处理的下一个指针
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
