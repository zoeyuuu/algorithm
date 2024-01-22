package linklist

// 142. 环形链表 II 2024-01-10 150
// https://leetcode.cn/problems/linked-list-cycle-ii/
// 与p141的区别：不仅要判断是否有环 还要判断环开始的初始位置

func Problem142() {
}

// 方法一 用map存储遍历后的结点
// 时间复杂度O(N) 空间复杂度O(N)
func detectCycle1(head *ListNode) *ListNode {
	mp := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := mp[head]; ok {
			return head
		}
		mp[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 方法二 快慢指针法 同p141
// 实现上注意：fast=slow时,不一定是环初始位置 经过数学推导 **让相遇节点和head同时移动 再次相遇时为环开始结点**
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	var met *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			met = fast
			for met != head {
				met = met.Next
				head = head.Next
			}
			return met
		}
	}
	return nil
}
