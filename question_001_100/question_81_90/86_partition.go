package question_81_90

// 86. 分隔链表
// https://leetcode-cn.com/problems/partition-list

func partition(head *ListNode, x int) *ListNode {
	var next, head1, tail1, head2, tail2 *ListNode
	next = head
	for next != nil {
		if next.Val < x {
			if head1 == nil {
				head1 = next
			} else {
				tail1.Next = next
			}
			tail1 = next
		} else {
			if head2 == nil {
				head2 = next
			} else {
				tail2.Next = next
			}
			tail2 = next
		}
		next = next.Next
	}
	if tail1 != nil {
		tail1.Next = head2
	}
	if tail2 != nil {
		tail2.Next = nil
	}
	if head1 != nil {
		return head1
	} else {
		return head2
	}
}
