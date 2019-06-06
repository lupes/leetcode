package question_91_100

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	start := &ListNode{Next: head}
	next := head
	var i = 1
	var hEnd = start
	var rStart = &ListNode{}
	var rEnd *ListNode
	for next != nil {
		t := next.Next
		if i >= m && i <= n {
			if i == m {
				rEnd = next
			}
			next.Next = rStart.Next
			rStart.Next = next
		} else if i < m {
			hEnd = next
		} else if i > n {
			break
		}
		next = t
		i++
	}
	rEnd.Next = next
	hEnd.Next = rStart.Next
	return start.Next
}
