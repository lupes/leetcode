package question_231_240

// 237. 删除链表中的节点
// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/comments/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}