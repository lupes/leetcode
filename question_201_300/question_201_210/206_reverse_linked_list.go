package question_201_210

// 206. 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/

func reverseList(head *ListNode) *ListNode {
	next := head
	head = &ListNode{}
	for next != nil {
		t := next.Next
		next.Next, next, head.Next = head.Next, t, next
	}
	return head.Next
}
