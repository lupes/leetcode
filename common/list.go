package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (self ListNode) String() string {
	if self.Next == nil {
		return fmt.Sprintf("{Val:%d}", self.Val)
	} else {
		return fmt.Sprintf("{Val:%d, Next:%s}", self.Val, self.Next)
	}
}

func NewList(nums ...int) *ListNode {
	var head = &ListNode{}
	var next = head
	for _, n := range nums {
		next.Next = &ListNode{Val: n}
		next = next.Next
	}
	return head.Next
}
