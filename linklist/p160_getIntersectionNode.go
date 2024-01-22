package linklist

// 160. 相交链表 easy 2023-12-11 178
// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/

func Problem160() {

}

// 思路：
// 链表相交->相交后一定相同
// 先求出两个链表的长度 然后设置快慢指针 让快指针先走长度差值
// 同时移动快慢指针 指针相同时即为相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	for curA != nil {
		lenA++
		curA = curA.Next
	}
	for curB != nil {
		lenB++
		curB = curB.Next
	}
	var step int
	var fast, slow *ListNode
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for step > 0 {
		fast = fast.Next
		step--
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
