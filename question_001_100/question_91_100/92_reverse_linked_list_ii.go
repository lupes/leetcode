package question_91_100

import (
	. "github.com/lupes/leetcode/common"
)

// 92. 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii
// Topics: 链表

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	head = &ListNode{Next: head}
	start := head
	for i := 1; i < m; i++ {
		start = start.Next
	}
	end, next := start.Next, start.Next
	for i := 0; i < n-m; i++ {
		next = end.Next
		start.Next, next.Next, end.Next = next, start.Next, end.Next.Next
	}

	return head.Next
}
