package question_421_430

// 430. 扁平化多级双向链表
// https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list
// Topics: 深度优先搜索 链表

type Node3 struct {
	Val   int
	Prev  *Node3
	Next  *Node3
	Child *Node3
}

func flatten(root *Node3) *Node3 {
	head, _ := flattenHelper(root)
	return head
}

func flattenHelper(root *Node3) (*Node3, *Node3) {
	var next, end = root, root
	for next != nil {
		if next.Child != nil {
			head, end := flattenHelper(next.Child)
			end.Next = next.Next
			if next.Next != nil {
				next.Next.Prev = end
			}
			next.Next = head
			head.Prev = next
			next.Child = nil
		}
		end = next
		next = next.Next
	}
	return root, end
}
