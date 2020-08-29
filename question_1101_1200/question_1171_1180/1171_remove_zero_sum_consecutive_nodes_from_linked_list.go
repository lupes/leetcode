package question_1171_1180

// 1171. 从链表中删去总和值为零的连续节点
// https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list
// Topics: 链表

import (
	. "github.com/lupes/leetcode/common"
)

func removeZeroSumSublists(head *ListNode) *ListNode {
	head = &ListNode{Val: 0, Next: head}
	for t := true; t; {
		t = false
		var sum, next, flag = 0, head, make(map[int]*ListNode)
		for next != nil {
			sum += next.Val
			if node, ok := flag[sum]; ok {
				node.Next, t = next.Next, true
				break
			}
			flag[sum], next = next, next.Next
		}
	}
	return head.Next
}
