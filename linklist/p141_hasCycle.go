package linklist

// 141. 环形链表 easy 2023-11-30 212
// https://leetcode.cn/problems/linked-list-cycle/description/

func Problem141() {

}

// 方法一
// 用map存储遍历后的结点
// 时间复杂度O(N) 空间复杂度O(N)
func hasCycle(head *ListNode) bool {
	mp := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := mp[head]; ok {
			return true
		}
		mp[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 方法二 龟兔赛跑算法 快慢指针
// 初始快慢指针都在head 每次快指针走2步 如果快指针能够追上慢指针 则链表中存在环
// 时间复杂度O(N) 空间复杂度O(1)
func hasCycle1(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
