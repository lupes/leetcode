package question_61_70

type ListNode struct {
	Val  int
	Next *ListNode
}

// 61. 旋转链表
// https://leetcode-cn.com/problems/rotate-list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	var next = head
	var tail *ListNode
	var size = 0
	for next != nil {
		size++
		tail = next
		next = next.Next
	}
	right := size - k%size
	if right == size {
		return head
	}
	next = head
	for i := 1; i < right; i++ {
		next = next.Next
	}
	tail.Next = head
	head = next.Next
	next.Next = nil
	return head
}
