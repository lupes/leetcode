package question_81_90

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. 删除排序链表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
// Topics: 链表

func deleteDuplicates(head *ListNode) *ListNode {
	next := head
	for next != nil {
		if next.Next != nil && next.Val == next.Next.Val {
			next.Next = next.Next.Next
		} else {
			next = next.Next
		}
	}
	return head
}
