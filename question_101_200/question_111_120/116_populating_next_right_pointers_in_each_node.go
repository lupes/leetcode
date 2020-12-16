package question_111_120

import (
	"fmt"
)

// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
// Topics: 树 深度优先搜索

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (self Node) String() string {
	return fmt.Sprintf("{Val:%d, Left:%s, Right:%s Next:%s}", self.Val, self.Left, self.Right, self.Next)
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	for next := root; next != nil; next = next.Next {
		next.Left.Next = next.Right
		if next.Next != nil {
			next.Right.Next = next.Next.Left
		}
	}
	connect(root.Left)
	return root
}
