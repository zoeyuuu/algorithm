package linklist

// 234. 回文链表 easy
// https://leetcode.cn/problems/palindrome-linked-list/description/

func isPalindrome(head *ListNode) bool {
	var nodeList []*ListNode
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	for i, j := 0, len(nodeList)-1; i < j; {
		if nodeList[i].Val != nodeList[j].Val {
			return false
		}
		i++
		j--
	}
	return true
}
