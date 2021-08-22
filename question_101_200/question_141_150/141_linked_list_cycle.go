package question_141_150

// 141. 环形链表
// https://leetcode-cn.com/problems/linked-list-cycle/
// Topics: 链表 双指针

import (
	. "github.com/lupes/leetcode/common"
)

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for next := head; next != nil; next = next.Next {
		if _, ok := m[next]; ok {
			return true
		} else {
			m[next] = struct{}{}
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	for low, fast := head, head.Next; fast != nil && fast.Next != nil; low, fast = low.Next, fast.Next.Next {
		if fast == low {
			return true
		}
	}
	return false
}
