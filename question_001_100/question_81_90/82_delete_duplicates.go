package question_81_90

func deleteDuplicates2(head *ListNode) *ListNode {
	var (
		next = head
		last *ListNode
		res  *ListNode
		n    *ListNode
		flag bool
	)

	for next != nil {
		for n = next.Next; n != nil; n = n.Next {
			if next.Val == n.Val {
				flag = true
			} else {
				break
			}
		}
		if !flag {
			if last == nil {
				last = next
				res = next
			} else {
				last.Next = next
				last = next
			}
			last.Next = nil
		}
		flag = false
		next = n
	}
	return res
}
