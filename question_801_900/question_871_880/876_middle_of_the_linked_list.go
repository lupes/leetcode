package question_871_880

// 876. 链表的中间结点
// https://leetcode-cn.com/problems/middle-of-the-linked-list
// Topics: 链表

import (
	. "github.com/lupes/leetcode/common"
)

func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
