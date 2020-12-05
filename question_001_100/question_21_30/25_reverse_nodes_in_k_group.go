package question_0011_0020

import (
	. "github.com/lupes/leetcode/common"
)

// 25. K 个一组翻转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group
// Topics: 链表

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	head = &ListNode{Next: head}
	var slow, fast = head, head.Next
	for i := 1; fast != nil; i++ {
		fast = fast.Next
		if i%k == 0 {
			tail, next := slow.Next, slow.Next.Next
			for j := 1; j < k; j++ {
				slow.Next, next.Next, next = next, slow.Next, next.Next
			}
			slow, tail.Next = tail, fast
		}
	}
	return head.Next
}
