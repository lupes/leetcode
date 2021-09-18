package question_1721_1730

// 1721. 交换链表中的节点
// https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list/
// Topics: 链表 双指针

import (
	. "github.com/lupes/leetcode/common"
)

func swapNodes(head *ListNode, k int) *ListNode {
	var last, fast, left = head, head, head
	for i := 1; i < k; i++ {
		fast = fast.Next
	}
	left = fast
	for fast.Next != nil {
		fast = fast.Next
		last = last.Next
	}
	left.Val, last.Val = last.Val, left.Val
	return head
}
