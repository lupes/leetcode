package question_0011_0020

// 21. 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	head := l1
	last := l1
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			last.Next = l2
			last = last.Next
			l2 = l2.Next
			last.Next = l1
		} else {
			last = l1
			l1 = l1.Next
		}
	}
	if l2 != nil {
		last.Next = l2
	}
	return head
}
