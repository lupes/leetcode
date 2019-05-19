package question_81_90

func deleteDuplicates2(head *ListNode) *ListNode {
	var (
		res  = &ListNode{Next: head}
		last = res
		now  = head
	)

	for now != nil {
		for now.Next != nil && now.Val == now.Next.Val {
			now = now.Next
		}
		if last.Next == now {
			last = last.Next
		} else {
			last.Next = now.Next
		}
		now = last.Next
	}
	return res.Next
}
