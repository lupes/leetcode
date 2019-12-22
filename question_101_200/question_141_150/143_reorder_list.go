package question_141_150

// 143. 重排链表
// https://leetcode-cn.com/problems/reorder-list
// Topics: 链表

import (
	. "github.com/lupes/leetcode/common"
)

func reorderList(head *ListNode) {
	var arr1, arr2 []*ListNode
	var next = head
	for next != nil {
		arr1 = append(arr1, next)
		next = next.Next
	}
	l := len(arr1)
	if l < 3 {
		return
	}
	for i := 0; i < l/2; i++ {
		arr2 = append(arr2, arr1[i], arr1[l-1-i])
	}
	if l%2 != 0 {
		arr2 = append(arr2, arr1[l/2])
	}
	next = arr2[0]
	for i := 1; i < len(arr2); i++ {
		next.Next = arr2[i]
		next = next.Next
	}
	next.Next = nil
}
