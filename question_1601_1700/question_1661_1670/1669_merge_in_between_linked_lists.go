package question_1661_1670

// 1669. 合并两个链表
// https://leetcode-cn.com/problems/merge-in-between-linked-lists/
// Topics: 链表

import (
	. "github.com/lupes/leetcode/common"
)

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var end1, end2 *ListNode
	for next := list2; next != nil; next = next.Next {
		end2 = next
	}
	for i, next := 0, list1; i <= b; i, next = i+1, next.Next {
		if i == a-1 {
			end1 = next
		} else if i == b {
			end1.Next = list2
			end2.Next = next.Next
		}
	}
	return list1
}
