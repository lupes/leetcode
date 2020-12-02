package common

import (
	"fmt"
	"math"
)

const Null = math.MaxInt32

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

func NewNodeV2(v ...int) *TreeNode {
	var root = &TreeNode{Val: v[0]}
	v = v[1:]
	var now = []*TreeNode{root}
	for {
		var next []*TreeNode
		for i := 0; i < len(now); i++ {
			if now[i] == nil {
				continue
			}
			if len(v) == 0 {
				return root
			}
			var left, right *TreeNode
			if v[0] != math.MaxInt32 {
				left = &TreeNode{Val: v[0]}
			}
			now[i].Left = left
			next = append(next, left)
			v = v[1:]

			if len(v) == 0 {
				return root
			}
			if v[0] != math.MaxInt32 {
				right = &TreeNode{Val: v[0]}
			}
			now[i].Right = right
			next = append(next, right)
			v = v[1:]
		}
		now = next
		if len(v) == 0 {
			return root
		}
	}
}
