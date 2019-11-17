package question_201_210

// 203. 移除链表元素
// https://leetcode-cn.com/problems/remove-linked-list-elements/
// Topics: 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{Next: head}
	next := head
	for next.Next != nil {
		if next.Next.Val == val {
			next.Next = next.Next.Next
		} else {
			next = next.Next
		}
	}
	return head.Next
}
