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
