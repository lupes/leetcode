package question_91_100

import . "github.com/lupes/leetcode/common"

// 92. 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	start := &ListNode{Next: head}
	next := head
	var i = 1
	var hEnd = start
	var reverseStart = &ListNode{}
	var reverseEnd *ListNode
	for next != nil {
		t := next.Next
		if i >= m && i <= n {
			if i == m {
				reverseEnd = next
			}
			next.Next = reverseStart.Next
			reverseStart.Next = next
		} else if i < m {
			hEnd = next
		} else if i > n {
			break
		}
		next = t
		i++
	}
	reverseEnd.Next = next
	hEnd.Next = reverseStart.Next
	return start.Next
}
