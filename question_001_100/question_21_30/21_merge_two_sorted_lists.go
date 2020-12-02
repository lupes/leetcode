package question_0011_0020

import (
	. "github.com/lupes/leetcode/common"
)

// 21. 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists
// Topics: 链表

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	next := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			next.Next, l2 = l2, nil
		} else if l2 == nil {
			next.Next, l1 = l1, nil
		} else if l1.Val > l2.Val {
			next.Next, l2 = l2, l2.Next
		} else {
			next.Next, l1 = l1, l1.Next
		}
		next = next.Next
	}
	return head.Next
}
