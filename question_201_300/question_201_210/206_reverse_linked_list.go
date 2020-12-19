package question_201_210

// 206. 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/
// Topics: 链表

func reverseList(head *ListNode) *ListNode {
	var res = &ListNode{}
	for next := head; next != nil; {
		next.Next, next, res.Next = res.Next, next.Next, next
	}
	return res.Next
}
