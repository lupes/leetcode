package question_11_20

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var nodes []*ListNode
	var tmp = head
	for tmp != nil {
		nodes = append(nodes, tmp)
		tmp = tmp.Next
	}
	size := len(nodes)
	if size < n {
		return nil
	} else if n == 1 {
		if size == 1 {
			return nil
		}
		nodes[size-2].Next = nil
	} else if size == n {
		head = head.Next
	} else {
		index := size - n
		nodes[index-1].Next = nodes[index+1]
	}
	return head
}
