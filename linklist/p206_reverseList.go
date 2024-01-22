package linklist

// 206. 反转链表 easy 2023-12-13 583
// https://leetcode.cn/problems/reverse-linked-list/description/
func Problem206() {

}

// 原地反转链表 用pre和cur双指针遍历
// 时间复杂度O(N) 空间复杂度O(1)
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		// 存储下一个要处理的位置
		next := cur.Next
		// 修改指向
		cur.Next = pre
		// 双指针分别后移
		pre = cur
		cur = next
	}
	// 循环结束pre在链表末尾 cur=nil
	return pre
}
