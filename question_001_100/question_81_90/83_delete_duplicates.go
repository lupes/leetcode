package question_81_90

type ListNode struct {
	Val  int
	Next *ListNode
}

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
