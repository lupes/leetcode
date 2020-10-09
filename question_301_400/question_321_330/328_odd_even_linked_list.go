package question_321_330

import (
	. "github.com/lupes/leetcode/common"
)

// 328. 奇偶链表
// https://leetcode-cn.com/problems/odd-even-linked-list
// Topics: 链表

func oddEvenList(head *ListNode) *ListNode {
	oddHead, evenHead := &ListNode{}, &ListNode{}
	oddEnd, evenEnd := oddHead, evenHead
	for i, next := 1, head; next != nil; i, next = i+1, next.Next {
		if i%2 == 1 {
			oddEnd.Next = next
			oddEnd = oddEnd.Next
		} else {
			evenEnd.Next = next
			evenEnd = evenEnd.Next
		}
	}
	oddEnd.Next, evenEnd.Next = evenHead.Next, nil
	return oddHead.Next
}
