package question_0011_0020

import (
	. "github.com/lupes/leetcode/common"
)

// 23. 合并K个排序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists
// Topics: 堆 链表 分治算法

func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoListsHelper(lists[0], lists[1])
	default:
		return mergeTwoListsHelper(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
	}
}

func mergeTwoListsHelper(l1 *ListNode, l2 *ListNode) *ListNode {
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
