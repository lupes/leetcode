package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (self *TreeNode) String() string {
	if self.Left != nil && self.Right != nil {
		return fmt.Sprintf("&TreeNode{Val:%d Left:%+v Right:%+v}", self.Val, self.Left, self.Right)
	} else if self.Left != nil {
		return fmt.Sprintf("&TreeNode{Val:%d Left:%+v}", self.Val, self.Left)
	} else if self.Right != nil {
		return fmt.Sprintf("&TreeNode{Val:%d Right:%+v}", self.Val, self.Right)
	} else {
		return fmt.Sprintf("&TreeNode{Val:%d}", self.Val)
	}
}

func NewNode(v int, son ...*TreeNode) *TreeNode {
	switch len(son) {
	case 0:
		return &TreeNode{Val: v}
	case 1:
		return &TreeNode{Val: v, Left: son[0]}
	default:
		return &TreeNode{Val: v, Left: son[0], Right: son[1]}
	}
}
