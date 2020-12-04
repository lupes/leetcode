package question_0011_0020

import (
	. "github.com/lupes/leetcode/common"
)

// 24. 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs
// Topics: 链表

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{Next: head}
	var slow, fast = head, head.Next
	for i := 1; fast != nil; i++ {
		fast = fast.Next
		if i%2 == 0 {
			slow, slow.Next, slow.Next.Next, slow.Next.Next.Next = slow.Next, slow.Next.Next, slow.Next.Next.Next, slow.Next
		}
	}
	return head.Next
}
