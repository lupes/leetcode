package question_0011_0020

// 24. 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs

func swapPairs(head *ListNode) *ListNode {
	next := head
	var one, two, res, last *ListNode
	one = head
	if one == nil {
		return nil
	}
	two = one.Next
	if two == nil {
		return head
	}
	res = two
	next = two.Next
	two.Next = one
	last = one
	for next != nil && next.Next != nil {
		one = next
		two = next.Next

		next = two.Next
		last.Next = two
		two.Next = one
		last = one
	}
	if next != nil {
		last.Next = next
	} else {
		last.Next = nil
	}
	return res
}
