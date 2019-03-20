package question_0011_0020

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	var (
		array = make([]*ListNode, k, k)
		res   *ListNode
		last  *ListNode
		next  = head
		index = 0
	)

	for next != nil {
		array[index] = next
		index++
		next = next.Next
		if index == k {
			for i := k - 1; i >= 0; i-- {
				if res == nil {
					res = array[i]
				} else {
					last.Next = array[i]
				}
				last = array[i]
			}
			index = 0
		}
	}
	for i := 0; i < index; i++ {
		if res == nil {
			res = array[i]
		} else {
			last.Next = array[i]
		}
		last = array[i]
	}
	last.Next = nil

	return res
}
