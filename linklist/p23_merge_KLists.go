package linklist

import "container/heap"

// 23 合并K个排序链表 hard
// https://leetcode.cn/problems/merge-k-sorted-lists/description/

func Problem23() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 5}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	l3 := &ListNode{Val: 1}
	l3.Next = &ListNode{Val: 4}

	lists := []*ListNode{l1, l2, l3}
	printList(mergeKLists2(lists))
}

// 2024.4.12 用堆实现
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	h := hp{}
	for _, v := range lists {
		if v != nil {
			h = append(h, v)
		}
	}
	heap.Init(&h)
	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

type hp []*ListNode

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}
func (h *hp) Push(x any) {
	*h = append(*h, x.(*ListNode))
}
func (h *hp) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

// 方法一 每次都两两合并 (暴力）
// 每个链表长度n k个链表
// 时间复杂度：第i次合并时间代价O((i+1)*n) 求和 最终为O(k*n^2)
// 空间复杂度O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} //注意空链表判断
	ans := lists[0]
	for i := 1; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i]) //mergeTwoLists：p21
	}
	return ans
}

// 方法二 类似于归并
// 时间复杂度 总共进行logk次归并 每次归并时间开销都是O(kn) 时间复杂度为O(knlogk)
// 空间复杂度 递归会使用到 O(logk)空间代价的栈空间
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}
func merge(lists []*ListNode, l, r int) *ListNode {
	//递归 参数是lists所有指针数组 和头尾lr
	if l == r {
		return lists[l] //如果相等返回指针 跟另一个指针进行mergeTwoLists
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r)) //mergeTwoLists p21
}
