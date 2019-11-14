package question_11_20

import "fmt"

// 19. 删除链表的倒数第N个节点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (self ListNode) String() string {
	return fmt.Sprintf("{Val:%d Next:%s}", self.Val, self.Next)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	head = &ListNode{Next: head}
	num := 0
	var next = head
	for next.Next != nil {
		num++
		next = next.Next
	}
	next = head
	for i := num - n; i > 0; i-- {
		next = next.Next
	}
	next.Next = next.Next.Next
	return head.Next
}
